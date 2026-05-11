package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	cleaned := strings.Fields(text)

	return cleaned
}
