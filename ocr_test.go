package dgocr_test

import (
	dgocr "github.com/darwinOrg/go-ppocr"
	"testing"
)

func TestOcrImageFile(t *testing.T) {
	dgocr.OcrImageFile("./imgs/11.jpg", "./result.json")
}
