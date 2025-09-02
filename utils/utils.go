package utils

import "fmt"

// ANSI color codes
const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue   = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan   = "\033[36m"
    ColorWhite  = "\033[37m"
    
    // Bright colors
    BrightRed    = "\033[91m"
    BrightGreen  = "\033[92m"
    BrightYellow = "\033[93m"
    BrightBlue   = "\033[94m"
    BrightPurple = "\033[95m"
    BrightCyan   = "\033[96m"
    BrightWhite  = "\033[97m"
    
    // Bold
    Bold = "\033[1m"
)

// Existing color functions (keep your current implementations if they exist)
func Red(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorRed, text, ColorReset)
}

func Green(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorGreen, text, ColorReset)
}

func Yellow(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorYellow, text, ColorReset)
}

func Blue(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorBlue, text, ColorReset)
}

func Purple(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorPurple, text, ColorReset)
}

func Cyan(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorCyan, text, ColorReset)
}

func White(text interface{}) string {
    return fmt.Sprintf("%s%v%s", ColorWhite, text, ColorReset)
}

// New bright color functions for highlighting
func BrightRed(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightRed, text, ColorReset)
}

func BrightGreen(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightGreen, text, ColorReset)
}

func BrightYellow(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightYellow, text, ColorReset)
}

func BrightBlue(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightBlue, text, ColorReset)
}

func BrightPurple(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightPurple, text, ColorReset)
}

func BrightCyan(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightCyan, text, ColorReset)
}

func BrightWhite(text interface{}) string {
    return fmt.Sprintf("%s%v%s", BrightWhite, text, ColorReset)
}

// Bold versions
func BoldRed(text interface{}) string {
    return fmt.Sprintf("%s%s%v%s", Bold, ColorRed, text, ColorReset)
}

func BoldGreen(text interface{}) string {
    return fmt.Sprintf("%s%s%v%s", Bold, ColorGreen, text, ColorReset)
}

func BoldYellow(text interface{}) string {
    return fmt.Sprintf("%s%s%v%s", Bold, ColorYellow, text, ColorReset)
}
