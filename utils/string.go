package utils

import (
	"regexp"
)

func TidyString(s string) string {
	regex := regexp.MustCompile(`[\n\t]`)
	cleanedString := regex.ReplaceAllString(s, "")
	return cleanedString
}
