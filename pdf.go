package dgocr

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgimk "github.com/darwinOrg/go-imagick"
	dglogger "github.com/darwinOrg/go-logger"
	"gopkg.in/gographics/imagick.v3/imagick"
	"strings"
	"time"
)

func AnnotateKeywordsForPdf(ctx *dgctx.DgContext, pdfFile string, keywords []string) (string, error) {
	start := time.Now().UnixMilli()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	pmw, err := dgimk.ConvertPdfToImage(ctx, mw, pdfFile)
	if err != nil {
		return "", err
	}
	defer pmw.Destroy()

	imageFile := pdfFile + ".jpg"
	if err := pmw.WriteImage(imageFile); err != nil {
		dglogger.Errorf(ctx, "WriteImage error：%v", err)
		return "", err
	}

	textRects, err := OcrImageFile(ctx, imageFile)
	if err != nil {
		return "", err
	}

	if len(textRects) == 0 {
		return imageFile, nil
	}

	cw := imagick.NewPixelWand()
	defer cw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	cw.SetColor("red")
	dw.SetStrokeColor(cw)

	cw.SetAlpha(0)
	dw.SetFillColor(cw)

	dw.SetStrokeWidth(1)
	dw.SetStrokeAntialias(true)

	for _, tr := range textRects {
		words := tr.Text
		wordsMetric := mw.QueryFontMetrics(dw, words)
		wordsWidth := wordsMetric.TextWidth
		leftTopX, leftTopY, rightBottomX, rightBottomY := tr.Rect.LeftTopX, tr.Rect.LeftTopY, tr.Rect.RightBottomX, tr.Rect.RightBottomY

		for _, keyword := range keywords {
			if strings.Contains(words, keyword) {
				keywordIndex := strings.Index(words, keyword)
				preWords := words[0:keywordIndex]

				var preWordsWidth float64
				if preWords != "" {
					preWordsMetric := mw.QueryFontMetrics(dw, preWords)
					preWordsWidth = preWordsMetric.TextWidth
				}

				keywordMetrics := mw.QueryFontMetrics(dw, keyword)
				keywordWidth := keywordMetrics.TextWidth

				dw.Rectangle(leftTopX+(rightBottomX-leftTopX)*(preWordsWidth/wordsWidth), leftTopY,
					leftTopX+(rightBottomX-leftTopX)*((preWordsWidth+keywordWidth)/wordsWidth), rightBottomY)
			}
		}
	}

	pmw.SetImageFormat("jpg")
	pmw.SetImageCompression(imagick.COMPRESSION_JPEG)

	if err := pmw.DrawImage(dw); err != nil {
		dglogger.Errorf(ctx, "DrawImage error：%v", err)
		return "", err
	}

	if err := pmw.WriteImage(imageFile); err != nil {
		dglogger.Errorf(ctx, "WriteImage error：%v", err)
		return "", err
	}

	cost := time.Now().UnixMilli() - start
	dglogger.Infof(ctx, "[file: %s] AnnotateKeywordsForPdf cost：%d ms", pdfFile, cost)

	return imageFile, nil
}
