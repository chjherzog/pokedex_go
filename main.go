package main

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	trimmedText := strings.Fields(text)
	var loweredText []string
	for i := range trimmedText {
		loweredText = append(loweredText, strings.ToLower(trimmedText[i]))
	}
	return loweredText
}

func main() {
	cleanedText := CleanInput("  hello    World")
	for i := range cleanedText {
		fmt.Println(cleanedText[i])
	}
	fmt.Println(cleanedText)
}
