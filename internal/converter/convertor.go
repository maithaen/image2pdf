package converter

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/maithaen/image2pdf/internal/config"
	"github.com/maithaen/image2pdf/pkg/utils"
)

type Converter struct {
    config *config.Config
    pdf    *gofpdf.Fpdf
}

func NewConverter(cfg *config.Config) *Converter {
    var pdf *gofpdf.Fpdf
    if cfg.UseA4 {
        pdf = gofpdf.New("P", "mm", "A4", "")
    } else {
        pdf = gofpdf.New("P", "mm", "", "")
    }

    return &Converter{
        config: cfg,
        pdf:    pdf,
    }
}

func (c *Converter) ConvertImages(images []string) error {
    if len(images) == 0 {
        utils.LogError("No images to convert")
        return fmt.Errorf("no images to convert")
    }

    tempFiles := make([]string, 0)
    defer func() {
        // Cleanup temporary files
        for _, f := range tempFiles {
            os.Remove(f)
        }
    }()

    for i, imgPath := range images {
        // Check if image needs rotation
        if c.config.AutoRotate {
            needsRotation, err := shouldRotate(imgPath)
            if err != nil {
                utils.LogError("Error checking rotation for %s: %v", imgPath, err)
                return fmt.Errorf("error checking rotation for %s: %v", imgPath, err)
            }

            if needsRotation {
                rotatedPath, err := rotateImage(imgPath)
                if err != nil {
                    utils.LogError("Error rotating image %s: %v", imgPath, err)
                    return fmt.Errorf("error rotating image %s: %v", imgPath, err)
                }
                tempFiles = append(tempFiles, rotatedPath)
                imgPath = rotatedPath
            }
        }

        // Add new page (except for first image)
        if i > 0 {
            c.pdf.AddPage()
        }

        // Get image dimensions
        opts := gofpdf.ImageOptions{
            ReadDpi: true,
        }

        imageInfo := c.pdf.RegisterImageOptions(imgPath, opts)
        if imageInfo == nil {
            return fmt.Errorf("error registering image: %s", imgPath)
        }

        // Calculate dimensions to fit page
        width, height := c.calculateDimensions(imageInfo.Width(), imageInfo.Height())

        // Add image to PDF
        c.pdf.AddPage()
        c.pdf.Image(imgPath, 0, 0, width, height, false, "", 0, "")
    }

    // Save the PDF
    return c.pdf.OutputFileAndClose(c.config.OutputFile)
}
func (c *Converter) calculateDimensions(imgWidth, imgHeight float64) (float64, float64) {
    // Get page dimensions
    pageWidth, pageHeight := c.pdf.GetPageSize()

    // Calculate scaling factors
    scaleX := pageWidth / imgWidth
    scaleY := pageHeight / imgHeight

    // Use the smaller scaling factor to maintain aspect ratio
    scale := scaleX
    if scaleY < scaleX {
        scale = scaleY
    }

    return imgWidth * scale, imgHeight * scale
}