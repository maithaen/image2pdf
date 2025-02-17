package converter

import (
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "os"
    "path/filepath"

    "github.com/disintegration/imaging"
    "github.com/maithaen/image2pdf/pkg/utils"
)

func shouldRotate(imgPath string) (bool, error) {
    file, err := os.Open(imgPath)
    if err != nil {
        return false, err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return false, err
    }

    bounds := img.Bounds()
    return bounds.Dx() > bounds.Dy(), nil
}

func rotateImage(imgPath string) (string, error) {
    utils.LogInfo("Rotating image: %s", filepath.Base(imgPath))

    // Load image with imaging package
    src, err := imaging.Open(imgPath)
    if err != nil {
        return "", err
    }

    // Rotate image 90 degrees clockwise
    rotated := imaging.Rotate90(src)

    // Create temporary file with original extension
    ext := filepath.Ext(imgPath)
    tempFile, err := os.CreateTemp("", "rotated-*"+ext)
    if err != nil {
        return "", err
    }
    tempPath := tempFile.Name()
    tempFile.Close()

    // Save rotated image with high quality
    err = imaging.Save(rotated, tempPath, imaging.JPEGQuality(100))
    if err != nil {
        os.Remove(tempPath)
        return "", err
    }

    utils.LogSuccess("Image rotated successfully")
    return tempPath, nil
}
