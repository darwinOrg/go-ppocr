package dgocr_test

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgocr "github.com/darwinOrg/go-ppocr"
	"gopkg.in/gographics/imagick.v3/imagick"
	"testing"
)

func TestAnnotateKeywordsForPdf(t *testing.T) {
	imagick.Initialize()
	defer imagick.Terminate()

	ctx := &dgctx.DgContext{TraceId: "123"}
	_, err := dgocr.AnnotateKeywordsForPdf(ctx, "1.pdf", []string{"Java", "MySQL", "15888888888"})
	if err != nil {
		return
	}
}
