package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/maithaen/image2pdf/internal/config"
	"github.com/maithaen/image2pdf/pkg/utils"
)

func init() {
    godotenv.Load()
}

const (
    exitSuccess = 0
    exitError   = 1
)

func main() {
    if err := run(); err != nil {
        utils.LogError("Error: %v", err)
        os.Exit(exitError)
    }
    os.Exit(exitSuccess)
}

func run() error {
    startTime := time.Now()
    log.SetFlags(log.Ldate | log.Ltime)

    cfg, err := config.ParseFlags()
    if err != nil {
        return fmt.Errorf("error parsing flags: %w", err)
    }

    if cfg.ShowVersion {
        utils.LogInfo("image2pdf version %s", config.Version)
        return nil
    }

    absInputDir, err := filepath.Abs(cfg.InputDir)
    if err != nil {
        return fmt.Errorf("error resolving input directory: %w", err)
    }
    cfg.InputDir = absInputDir

    absOutputFile, err := filepath.Abs(cfg.OutputFile)
    if err != nil {
        return fmt.Errorf("error resolving output file: %w", err)
    }
    cfg.OutputFile = absOutputFile

    utils.LogInfo("Scanning directory: %s", cfg.InputDir)
    
    // Initialize PDF converter
    pdfConverter := utils.NewPDFConverter(cfg.UseA4, cfg.AutoRotate)
    pdfConverter.InitDocument()

    imageCount := 0
    err = filepath.Walk(cfg.InputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil || info.IsDir() {
            return err
        }

        if utils.IsImageFile(path) {
            if cfg.JpegOnly && !strings.HasSuffix(strings.ToLower(path), ".jpg") && 
               !strings.HasSuffix(strings.ToLower(path), ".jpeg") {
                return nil
            }

            utils.LogInfo("Processing: %s", filepath.Base(path))
            img, err := utils.LoadImage(path)
            if err != nil {
                utils.LogWarning("Skipping %s: %v", filepath.Base(path), err)
                return nil
            }

            if err := pdfConverter.AddImagePage(img); err != nil {
                utils.LogWarning("Failed to add %s: %v", filepath.Base(path), err)
                return nil
            }
            imageCount++
        }
        return nil
    })

    if err != nil {
        return fmt.Errorf("error scanning directory: %w", err)
    }

    if imageCount == 0 {
        return fmt.Errorf("no valid images found in directory: %s", cfg.InputDir)
    }

    utils.LogSuccess("Found %d images", imageCount)

    if err := pdfConverter.SavePDF(cfg.OutputFile); err != nil {
        return fmt.Errorf("error saving PDF: %w", err)
    }

    duration := time.Since(startTime)
    utils.LogSuccess("Conversion completed in %v", duration)
    utils.LogSuccess("PDF saved to: %s", cfg.OutputFile)

    return nil
}
