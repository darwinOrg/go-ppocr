package dgocr

import (
	"encoding/json"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"os"
	"os/exec"
	"time"
)

type Rect struct {
	LeftTopX     float64 `json:"leftTopX"`
	LeftTopY     float64 `json:"leftTopY"`
	RightBottomX float64 `json:"rightBottomX"`
	RightBottomY float64 `json:"rightBottomY"`
}

type TextRect struct {
	Text string `json:"text"`
	Rect *Rect  `json:"rect"`
}

func OcrImageFile(ctx *dgctx.DgContext, sourceImageFile string) ([]*TextRect, error) {
	start := time.Now().UnixMilli()
	destOcrFile := sourceImageFile + ".ocr"
	cmd := exec.Command("python", "ocr.py", sourceImageFile, destOcrFile)
	output, err := cmd.Output()
	if err != nil {
		dglogger.Errorf(ctx, "exec python ocr.py error: %v", err)
		return nil, err
	}
	defer func() {
		os.Remove(destOcrFile)
	}()
	dglogger.Debugf(ctx, "paddleocr output: %s", string(output))

	var d [][][][]any
	err = json.Unmarshal(output, &d)
	if err != nil {
		dglogger.Errorf(ctx, "json[%s] unmarshal error: %v", string(output), err)
		return nil, err
	}

	var textRects []*TextRect
	if len(d) > 0 {
		for _, d1 := range d {
			if len(d1) > 0 {
				for _, d2 := range d1 {
					leftTop := d2[0][0].([]any)
					rightBottom := d2[0][2].([]any)

					rect := &Rect{
						LeftTopX:     leftTop[0].(float64),
						LeftTopY:     leftTop[1].(float64),
						RightBottomX: rightBottom[0].(float64),
						RightBottomY: rightBottom[1].(float64),
					}

					text := d2[1][0].(string)

					textRects = append(textRects, &TextRect{
						Text: text,
						Rect: rect,
					})
				}
			}
		}
	}

	cost := time.Now().UnixMilli() - start
	dglogger.Infof(ctx, "[file: %s] OcrImageFile cost：%d ms", sourceImageFile, cost)

	return textRects, nil
}
