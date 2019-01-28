package chapter01

import (
	"strings"
)

// Trims the whitespace from the input string, and replaces the internal spaces
// with `"%20"`.
func URLIfy(someString string) string {
	someString = strings.TrimSpace(someString)

	return strings.Replace(someString, " ", "%20", -1)
}
