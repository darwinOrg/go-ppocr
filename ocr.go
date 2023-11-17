package dgocr

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"os/exec"
)

func OcrImageFile(ctx *dgctx.DgContext, sourceImageFile string, destJsonFile string) error {
	cmd := exec.Command("python", "ocr.py", sourceImageFile, destJsonFile)
	output, err := cmd.Output()
	if err != nil {
		dglogger.Errorf(ctx, "exec python ocr.py error: %v", err)
		return err
	}
	dglogger.Debugf(ctx, "output: %s", string(output))
	return nil
}
