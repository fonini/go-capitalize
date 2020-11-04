// Package capitalize implements capitalization of proper names
//
// As a simple example:
//
// 	name, _ := capitalize.Capitalize("jonnas fonini")
// 	fmt.Print(name) // will output: "Jonnas Fonini"
//
package capitalize

import (
	"regexp"
	"strings"
)

var exceptions []string = []string{
	"de", "di", "do", "da", "dos", "das", "dello", "della", "dalla", "dal", "del",
	"la", "e", "em", "na", "no", "nas", "nos", "van", "von", "y",
}

// Options is a configuration struct. It allows to add new exceptions
type Options struct {
	Exceptions []string
}

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexRomanNumeral = regexp.MustCompile(`^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
var point = regexp.MustCompile(`/\./g`)

// Capitalize returns the capitalized string.
func Capitalize(str string, options ...Options) (string, error) {
	if len(options) > 0 {
		exceptions = append(exceptions, options[0].Exceptions...)
	}

	var pointSpace = `. `
	var space = ` `

	// Replace all points with point+space
	str = point.ReplaceAllString(str, pointSpace)

	// Remove redundant spaces
	str = regexMultipleSpaces.ReplaceAllString(str, space)

	// Split the words by space
	split := strings.Split(str, space)

	parts := []string{}

	// Capitalize each word
	for i := range split {
		parts = append(parts, strings.Title(strings.ToLower(split[i])))
	}

	for i := 0; i < len(parts); i++ {
		// Find and replace the words in the exceptions list
		for j := 0; j < len(exceptions); j++ {
			if strings.ToLower(parts[i]) == strings.ToLower(exceptions[j]) {
				parts[i] = exceptions[j]
			}
		}

		// Capitalize all the chars if the part is a Roman Number
		if regexRomanNumeral.MatchString(strings.ToUpper(parts[i])) {
			parts[i] = strings.ToUpper(parts[i])
		}
	}

	// Join the pieces
	return strings.Join(parts, space), nil
}
