package converter

import (
	"fmt"
	"image"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
	_ "golang.org/x/image/webp"
)

func ConvertImage(inputPath, outputPath, format string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	format = strings.ToLower(format)
	var errSave error

	switch format {
	case "png":
		errSave = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(3))
	case "jpg", "jpeg":
		errSave = imaging.Save(img, outputPath, imaging.JPEGQuality(95))
	case "webp":
		errSave = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(3)) // Adjust as needed
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	if errSave != nil {
		return fmt.Errorf("failed to save image: %v", errSave)
	}

	fmt.Println("Conversion successful:", outputPath)
	return nil
}
