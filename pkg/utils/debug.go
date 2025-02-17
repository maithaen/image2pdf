
package utils

import (
	"fmt"
)

const (
	// Color codes for terminal output
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

// LogSuccess prints a success message in green
func LogSuccess(format string, a ...interface{}) {
	fmt.Printf(ColorGreen+"[SUCCESS] "+format+ColorReset+"\n", a...)
}

// LogError prints an error message in red
func LogError(format string, a ...interface{}) {
	fmt.Printf(ColorRed+"[ERROR] "+format+ColorReset+"\n", a...)
}

// LogWarning prints a warning message in yellow
func LogWarning(format string, a ...interface{}) {
	fmt.Printf(ColorYellow+"[WARNING] "+format+ColorReset+"\n", a...)
}

// LogInfo prints an info message in blue
func LogInfo(format string, a ...interface{}) {
	fmt.Printf(ColorBlue+"[INFO] "+format+ColorReset+"\n", a...)
}

// LogDebug prints a debug message in cyan
func LogDebug(format string, a ...interface{}) {
	fmt.Printf(ColorCyan+"[DEBUG] "+format+ColorReset+"\n", a...)
}
