package validator

import (
	"path/filepath"
	"strings"

	"github.com/maithaen/image2pdf/pkg/utils"
)

var (
    SupportedExtensions = map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
        ".png":  true,
        ".gif":  true,
    }

    JpegExtensions = map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
    }
)

func IsValidImage(filename string, jpegOnly bool) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    utils.LogDebug("Checking file: %s with extension: %s", filename, ext)
    if jpegOnly {
        return JpegExtensions[ext]
    }
    return SupportedExtensions[ext]
}