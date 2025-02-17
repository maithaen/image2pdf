package utils

import (
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
)

type PDFConverter struct {
	pdf        *gofpdf.Fpdf
	useA4      bool
	autoRotate bool
}

func NewPDFConverter(useA4 bool, autoRotate bool) *PDFConverter {
	return &PDFConverter{
		useA4:      useA4,
		autoRotate: autoRotate,
	}
}

func (p *PDFConverter) InitDocument() {
	if p.useA4 {
		p.pdf = gofpdf.New("P", "mm", "A4", "")
	} else {
		p.pdf = gofpdf.New("P", "mm", "", "")
	}
}

func (p *PDFConverter) AddImagePage(img image.Image) error {
	// Convert image to JPEG in memory
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 95}); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	// Create temporary image file in PDF
	imgID := fmt.Sprintf("img%d", p.pdf.PageCount()+1)
	p.pdf.RegisterImageOptionsReader(imgID, gofpdf.ImageOptions{ImageType: "JPEG"}, buf)

	// Get image dimensions
	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	// Calculate scaling to fit page
	pageWidth, pageHeight := p.pdf.GetPageSize()
	scale := pageWidth / imgWidth
	if imgHeight*scale > pageHeight {
		scale = pageHeight / imgHeight
	}

	// Calculate centered position
	width := imgWidth * scale
	height := imgHeight * scale
	x := (pageWidth - width) / 2
	y := (pageHeight - height) / 2

	// Add new page and image
	p.pdf.AddPage()
	p.pdf.Image(imgID, x, y, width, height, false, "", 0, "")

	return nil
}

func (p *PDFConverter) SavePDF(outputPath string) error {
	return p.pdf.OutputFileAndClose(outputPath)
}
