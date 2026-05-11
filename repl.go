package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	cleaned := strings.Fields(text)

	return cleaned
}
