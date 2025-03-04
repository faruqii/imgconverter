package converter

import (
	"bytes"
	"fmt"
	"os"

	"github.com/disintegration/imaging"
	"github.com/jdeng/goheif"
)

// ConvertHEIC converts a HEIC file to the desired format (PNG, JPG, WebP)
func ConvertHEIC(inputPath, outputPath, format string) error {
	// Read the HEIC file
	file, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Decode HEIC
	img, err := goheif.Decode(bytes.NewReader(file))
	if err != nil {
		return fmt.Errorf("failed to decode HEIC: %v", err)
	}

	// Convert to desired format
	switch format {
	case "png":
		err = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(3))
	case "jpg", "jpeg":
		err = imaging.Save(img, outputPath, imaging.JPEGQuality(95))
	case "webp":
		err = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(3))
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return fmt.Errorf("failed to save converted image: %v", err)
	}

	fmt.Println("HEIC Conversion successful:", outputPath)
	return nil
}
