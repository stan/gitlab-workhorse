package artifacts

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"syscall"

	"gitlab.com/gitlab-org/gitlab-workhorse/internal/api"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/filestore"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/helper"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/upload"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/zipartifacts"
)

type artifactsUploadProcessor struct {
	opts   *filestore.SaveFileOpts
	stored bool
}

func (a *artifactsUploadProcessor) generateMetadataFromZip(ctx context.Context, file *filestore.FileHandler) (*filestore.FileHandler, error) {
	metaReader, metaWriter := io.Pipe()
	defer metaWriter.Close()

	metaOpts := &filestore.SaveFileOpts{
		LocalTempPath:  a.opts.LocalTempPath,
		TempFilePrefix: "metadata.gz",
	}
	if metaOpts.LocalTempPath == "" {
		metaOpts.LocalTempPath = os.TempDir()
	}

	fileName := file.LocalPath
	if fileName == "" {
		fileName = file.RemoteURL
	}

	zipMd := exec.CommandContext(ctx, "gitlab-zip-metadata", fileName)
	zipMd.Stderr = os.Stderr
	zipMd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	zipMd.Stdout = metaWriter

	if err := zipMd.Start(); err != nil {
		return nil, err
	}
	defer helper.CleanUpProcessGroup(zipMd)

	type saveResult struct {
		error
		*filestore.FileHandler
	}
	done := make(chan saveResult)
	go func() {
		var result saveResult
		result.FileHandler, result.error = filestore.SaveFileFromReader(ctx, metaReader, -1, metaOpts)

		done <- result
	}()

	if err := zipMd.Wait(); err != nil {
		if st, ok := helper.ExitStatus(err); ok && st == zipartifacts.StatusNotZip {
			return nil, nil
		}
		return nil, err
	}

	metaWriter.Close()
	result := <-done
	return result.FileHandler, result.error
}

func (a *artifactsUploadProcessor) ProcessFile(ctx context.Context, formName string, file *filestore.FileHandler, writer *multipart.Writer) error {
	//  ProcessFile for artifacts requires file form-data field name to eq `file`

	if formName != "file" {
		return fmt.Errorf("invalid form field: %q", formName)
	}
	if a.stored {
		return fmt.Errorf("artifacts request contains more than one file")
	}
	a.stored = true

	select {
	case <-ctx.Done():
		return fmt.Errorf("ProcessFile: context done")

	default:
		// TODO: can we rely on disk for shipping metadata? Not if we split workhorse and rails in 2 different PODs
		metadata, err := a.generateMetadataFromZip(ctx, file)
		if err != nil {
			return fmt.Errorf("generateMetadataFromZip: %v", err)
		}

		if metadata != nil {
			for k, v := range metadata.GitLabFinalizeFields("metadata") {
				writer.WriteField(k, v)
			}
		}
	}
	return nil
}

func (a *artifactsUploadProcessor) ProcessField(ctx context.Context, formName string, writer *multipart.Writer) error {
	return nil
}

func (a *artifactsUploadProcessor) Finalize(ctx context.Context) error {
	return nil
}

func (a *artifactsUploadProcessor) Name() string {
	return "artifacts"
}

func UploadArtifacts(myAPI *api.API, h http.Handler) http.Handler {
	return myAPI.PreAuthorizeHandler(func(w http.ResponseWriter, r *http.Request, a *api.Response) {
		mg := &artifactsUploadProcessor{opts: filestore.GetOpts(a)}

		upload.HandleFileUploads(w, r, h, a, mg)
	}, "/authorize")
}
