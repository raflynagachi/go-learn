package main

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		switch v {
		case '❗':
			return "recommendation"
		default:
			return "default"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for idx, v := range runes {
		if v == oldRune {
			runes[idx] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) >= limit
}

func main() {
	fmt.Println("❗ recommended product")
	fmt.Println(Application("❗ recommended product"))
	fmt.Println(Replace("❗ recommended product", '❗', '?'))
	fmt.Println(WithinLimit("❗ recommended product", 20))
}
