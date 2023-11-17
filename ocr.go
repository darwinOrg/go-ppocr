package dgocr

import (
	"fmt"
	"os/exec"
)

func OcrImageFile(sourceImageFile string, destJsonFile string) error {
	cmd := exec.Command("python", "ocr.py", sourceImageFile, destJsonFile)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", string(output))
	}

	return err
}
