// Package capitalize implements capitalization of proper names
//
// As a simple example:
//
// 	name, _ := capitalize.Capitalize("jonnas fonini")
// 	fmt.Print(name) // will output: "Jonnas Fonini"
//
// With additional exceptions
// 	options := capitalize.Options{
// 		Exceptions: []string{"of"},
// 	}
//
// 	name, _ = capitalize.Capitalize("gørvel fadersdotter of giske", options)
// 	fmt.Println(name) // will output: "Gørvel Fadersdotter of Giske"
//
// 	options = capitalize.Options{
// 		Exceptions: []string{"McElroy"},
// 	}
//
// 	name, _ = capitalize.Capitalize("john mcelroy", options)
// 	fmt.Println(name) // will output: "John McElroy"
//
package capitalize

import (
	"regexp"
	"strings"
)

var exceptions []string = []string{
	"de", "di", "do", "da", "dos", "das", "dello", "della", "dalla", "dal", "del",
	"la", "e", "em", "na", "no", "nas", "nos", "the", "van", "von", "y",
}

var surnames []string = []string{
	"McCain",
}

// Options is a configuration struct.
type Options struct {
	// Exceptions is a list of words that must remain in lower case. e.g. von, van.
	// https://en.wikipedia.org/wiki/Nobiliary_particle
	Exceptions []string
	// Surnames is a list of surnames that have a special capitalization,
	// so they will remain exactly as in the list. e.g. McCain, McElroy
	Surnames []string
}

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexRomanNumeral = regexp.MustCompile(`^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
var point = regexp.MustCompile(`/\./g`)

// Capitalize returns the capitalized string.
// You can pass along a Options struct to add more exceptions
func Capitalize(str string, options ...Options) (string, error) {
	if len(options) > 0 {
		exceptions = append(exceptions, options[0].Exceptions...)
		surnames = append(surnames, options[0].Surnames...)
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

		// Find and replace the words in the surnames list
		for j := 0; j < len(surnames); j++ {
			if strings.ToLower(parts[i]) == strings.ToLower(surnames[j]) {
				parts[i] = surnames[j]
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
