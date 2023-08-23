package utils

import "strings"

// remove new line
func RemoveNewLine(str string) string {
	return strings.ReplaceAll(str, "\n", "")
}
