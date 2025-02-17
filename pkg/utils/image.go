package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// SupportedImageTypes returns a list of supported image extensions
func SupportedImageTypes() []string {
	return []string{".jpg", ".jpeg", ".png"}
}

// IsImageFile checks if the given file is a supported image
func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, supportedExt := range SupportedImageTypes() {
		if ext == supportedExt {
			return true
		}
	}
	return false
}

// LoadImage loads an image from file path
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var img image.Image
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", ext)
	}

	return img, err
}
