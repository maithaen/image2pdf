package scanner

import (
    "os"
    "path/filepath"

    "github.com/maithaen/image2pdf/internal/config"
    "github.com/maithaen/image2pdf/internal/validator"
)

type Scanner struct {
    config *config.Config
    files  []string
}

func NewScanner(cfg *config.Config) *Scanner {
    return &Scanner{
        config: cfg,
        files:  make([]string, 0),
    }
}

func (s *Scanner) ScanDirectory() ([]string, error) {
    err := filepath.Walk(s.config.InputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            // Calculate current depth
            relPath, err := filepath.Rel(s.config.InputDir, path)
            if err != nil {
                return err
            }

            depth := len(filepath.SplitList(relPath))
            if depth >= s.config.ScanLevel {
                return filepath.SkipDir
            }
            return nil
        }

        // Check if file is a valid image
        if validator.IsValidImage(path, s.config.JpegOnly) {
            s.files = append(s.files, path)
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    return s.files, nil
}