package dgocr_test

import (
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/darwinOrg/go-common/utils"
	dglogger "github.com/darwinOrg/go-logger"
	dgocr "github.com/darwinOrg/go-ppocr"
	"os"
	"testing"
)

func TestOcrImageFile(t *testing.T) {
	sourceImageFile := os.Getenv("imageFile")
	ctx := &dgctx.DgContext{TraceId: "123"}
	textRects, err := dgocr.OcrImageFile(ctx, sourceImageFile)
	if err != nil {
		dglogger.Error(ctx, err)
		return
	} else {
		dglogger.Infof(ctx, "%s", utils.MustConvertBeanToJsonString(textRects))
	}
}
