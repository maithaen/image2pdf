# Image2PDF Converter command-line tool

A simple and efficient command-line tool to convert images into PDF files. Supports multiple image formats and provides flexible configuration options.

## Features

- Convert multiple images to a single PDF
- Support for various image formats (JPG, JPEG, PNG, GIF)
- Option to process only JPG/JPEG files
- Customizable output filename
- A4 format support
- Recursive directory scanning with configurable depth
- Automatic image rotation support

## Installation

1. Ensure you have Go installed on your system
2. Clone the repository:

git clone <https://github.com/maithaen/image2pdf.git>
cd image2pdf

## Build

go build -o image2pdf.exe main.go

## Usage

./image2pdf [options]

## Command Line Options

-all: Convert all supported image types (jpg, jpeg, png, gif)
-jpg: Convert only jpg/jpeg images
-o: Specify output PDF filename (default: "output.pdf")
-a4: Use A4 paper size
-dir: Specify input directory to scan for images (default: current directory)
-l: Directory scan level (1: root only, 2: one level deep, 3: two levels deep)
-r: Rotate landscape images to portrait orientation

## Examples

Convert all images in current directory:

./image2pdf.exe -all

Convert only JPG files in a specific directory:

./image2pdf.exe -jpg -dir /path/to/images -o my_photos.pdf

Create A4-sized PDF with recursive scanning:

./image2pdf.exe -a4 -dir /path/to/images -l 2 -o my_photos.pdf

Convert images with automatic rotation:

./image2pdf.exe -all -r -dir /path/to/images -o my_photos.pdf

## Dependencies

- github.com/jung-kurt/gofpdf - PDF generation library

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
