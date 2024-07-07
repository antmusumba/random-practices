package main

import (
    "fmt"
    "golang.org/x/term"
    "os"
    "strings"
)

// GetTerminalWidth returns the current width of the terminal in characters.
func getTerminalWidth() int {
    width, _, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        fmt.Println("Error getting terminal size:", err)
        return 80 // Default width
    }
    return width
}

// CenterAlign centers the text based on the terminal width.
func centerAlign(text string) {
    width := getTerminalWidth()
    padding := (width - len(text)) / 2
    if padding < 0 {
        padding = 0 // Prevent negative paddingk
    }
    fmt.Printf("%s%s\n", strings.Repeat(" ", padding), text)
}

// LeftAlign aligns the text to the left based on the terminal width.
func leftAlign(text string) {
    fmt.Println(text)
}

// RightAlign aligns the text to the right based on the terminal width.
func rightAlign(text string) {
    width := getTerminalWidth()
    padding := width - len(text)
    if padding < 0 {
        padding = 0 // Prevent negative padding
    }
    fmt.Printf("%s%s\n", strings.Repeat(" ", padding), text)
}

// WrapText wraps the text to fit within the terminal width.
func wrapText(text string) []string {
    width := getTerminalWidth()
    var lines []string
    for len(text) > width {
        splitAt := strings.LastIndex(text[:width], " ")
        if splitAt == -1 { // No space found, force split
            splitAt = width
        }
        lines = append(lines, text[:splitAt])
        text = text[splitAt:]
        if len(text) > 0 && text[0] == ' ' {
            text = text[1:] // Remove leading space
        }
    }
    lines = append(lines, text)
    return lines
}

func main() {
    text := "This is a very"

    fmt.Println("Center Aligned:")
    for _, line := range wrapText(text) {
        centerAlign(line)
    }

    fmt.Println("\nLeft Aligned:")
    for _, line := range wrapText(text) {
        leftAlign(line)
    }

    fmt.Println("\nRight Aligned:")
    for _, line := range wrapText(text) {
        rightAlign(line)
    }
}
