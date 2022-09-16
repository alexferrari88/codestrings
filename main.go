package codestrings

import "strings"

func populateMap(m map[string]struct{}, s []string) {
	for _, v := range s {
		m[v] = struct{}{}
	}
}

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
			if started {
				if strings.TrimSpace(currentString) != "" && c == delimiter && source[i-1] != '\\' {
					// ignore empty strings or being made of only spaces
					// c == delimiter is to avoid cases like "it's a beautiful day"
					stringsList = append(stringsList, currentString)
					started = false
					currentString = ""
					delimiter = 0
					continue
				} else if c != delimiter {
					// c != delimiter is to avoid cases like "it's a beautiful day"
					currentString += string(c)
					continue
				}
			} else if !started && i > 0 && source[i-1] == '\\' {
				// ignore escaped strings
				continue
			} else if !started && i != len(source)-1 {
				// ignore unterminated strings
				started = true
				delimiter = c
				continue
			}
		}
		if started && c != '\\' && i != len(source)-1 {
			currentString += string(c)
		}
	}
	return stringsList
}
