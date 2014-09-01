// Package stringutils provides some higher-level string operations that aren't provided by the strings package.
package stringutils

import (
	"strings"
	"unicode"
)

// UcFirst returns a string that matches the input string, but with the first letter capitalized.
func UcFirst(input string) string {
	if len(input) == 0 {
		return ""
	}

	return strings.ToUpper(input[0:1]) + input[1:]
}

// LcFirst returns a string that matches the input string, but with the first letter converted to lowercase.
func LcFirst(input string) string {
	if len(input) == 0 {
		return ""
	}

	return strings.ToUpper(input[0:1]) + input[1:]
}

// CamelToSpinal will convert a string from camelCase or UpperCamelCase into spinal-case.
func CamelToSpinal(input string) string {
	// Ensure the input is in UpperCamelCase
	input = UcFirst(input)
	
	var words []string

	lastCap := 0
	for i, runeVal := range input {
		if i > 0 && unicode.IsUpper(runeVal) {
			words = append(words, strings.ToLower(input[lastCap:i]))
			lastCap = i
		}
	}

	if len(input) > lastCap {
		words = append(words, strings.ToLower(input[lastCap:]))
	}

	return strings.Join(words, "-")
}

// SpinalToCamel will convert a string from spinal-case or Train-Case into UpperCamelCase
func SpinalToCamel(input string) string {
	// Ensure the input is all lowercase
	input = strings.ToLower(input)
	
	words := strings.Split(input, "-")

	for i, word := range words {
		words[i] = UcFirst(word)
	}

	return strings.Join(words, "")
}
