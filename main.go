package codestrings

import "strings"

func populateMap(m map[string]struct{}, s []string) {
	for _, v := range s {
		m[v] = struct{}{}
	}
}

// ExtractStrings extracts all strings from a source code
// stringDelimiters is a list of delimiters that can be used to start and end a string
// For instance, in JavaScript, strings can be delimited by " or ', e.g. "hello" or 'world'
// If stringDelimiters is empty, it will use the default ones: " and '
// It will ignore escaped strings and unterminated strings
func ExtractStrings(source string, stringDelimiters []string) []string {
	var stringsList []string
	var started bool
	var delimiter byte
	var currentString string
	if len(stringDelimiters) == 0 {
		stringDelimiters = []string{"\"", "'"}
	}
	// Create a map for string delimiters for faster lookup
	// A string delimiter is the character that starts and ends a string
	symbolsMap := make(map[string]struct{}, len(stringDelimiters))
	populateMap(symbolsMap, stringDelimiters)
	for i := range source {
		c := source[i]
		if _, ok := symbolsMap[string(c)]; ok {
			if started && strings.TrimSpace(currentString) != "" {
				if c == delimiter && source[i-1] != '\\' {
					// we found the end of a string and it's not empty
					// source[i-1] != '\\' means that the delimiter is not escaped so it's an actual string delimiter
					stringsList = append(stringsList, currentString)
					started = false
					currentString = ""
					delimiter = 0
					continue
				} else if c != delimiter {
					// if the delimiter is not the same as the one that started the string, it's not a string delimiter
					// e.g. var a = "it's a beautiful day"
					currentString += string(c)
					continue
				}
			} else if !started && i != len(source)-1 {
				// we found the start of a string
				// we also check that the string delimiter is not at the end of the source code
				// so we ignore unterminated strings
				started = true
				delimiter = c
				continue
			}
		}
		if started && i != len(source)-1 {
			currentString += string(c)
		}
	}
	return stringsList
}
