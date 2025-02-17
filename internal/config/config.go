package config

import (
	"flag"
	"fmt"

	"github.com/maithaen/image2pdf/pkg/utils"
)

const Version = "0.0.1"

type Config struct {
    ConvertAll  bool
    JpegOnly    bool
    OutputFile  string
    UseA4      bool
    InputDir   string
    ScanLevel  int
    AutoRotate bool
    ShowVersion bool 
}

func ParseFlags() (*Config, error) {
    cfg := &Config{}

    flag.BoolVar(&cfg.ConvertAll, "all", false, "Convert all supported image types")
    flag.BoolVar(&cfg.JpegOnly, "jpg", false, "Convert only jpg/jpeg images")
    flag.StringVar(&cfg.OutputFile, "o", "output.pdf", "Output PDF filename")
    flag.BoolVar(&cfg.UseA4, "a4", false, "Use A4 paper size")
    flag.StringVar(&cfg.InputDir, "dir", ".", "Input directory to scan")
    flag.IntVar(&cfg.ScanLevel, "l", 1, "Directory scan level (1: root only, 2: one level deep, 3: two levels deep)")
    flag.BoolVar(&cfg.AutoRotate, "r", false, "Rotate landscape images to portrait")
    flag.BoolVar(&cfg.ShowVersion, "version", false, "Show version information")

    flag.Parse()

    utils.LogInfo("Parsing configuration flags...")

    // Validate configuration
    if err := cfg.validate(); err != nil {
        return nil, err
    }

    utils.LogSuccess("Configuration validated successfully")
    return cfg, nil
}

func (c *Config) validate() error {
    if !c.ConvertAll && !c.JpegOnly {
        utils.LogWarning("No conversion type specified, defaulting to all image types")
        c.ConvertAll = true
    }

    if c.ScanLevel < 1 || c.ScanLevel > 3 {
        utils.LogError("Invalid scan level: %d (must be between 1 and 3)", c.ScanLevel)
        return fmt.Errorf("scan level must be between 1 and 3")
    }

    if c.AutoRotate {
        utils.LogInfo("Auto-rotation enabled - landscape images will be rotated to portrait")
    }

    return nil
}
