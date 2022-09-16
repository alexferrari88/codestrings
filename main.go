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
	var currentString string
	if len(stringDelimiters) == 0 {
		stringDelimiters = []string{"\"", "'"}
	}
	symbolsMap := make(map[string]struct{}, len(stringDelimiters))
	populateMap(symbolsMap, stringDelimiters)
	for i := range source {
		c := source[i]
		if _, ok := symbolsMap[string(c)]; ok {
			if started {
				if strings.TrimSpace(currentString) != "" { // ignore empty strings or with only spaces
					stringsList = append(stringsList, currentString)
				}
				started = false
				currentString = ""
				continue
			} else if !started && i > 0 && source[i-1] == '\\' { // ignore escaped strings
				continue
			} else if !started && i != len(source)-1 { // ignore unterminated strings
				started = true
				continue
			}
		}
		if started {
			currentString += string(c)
		}
	}
	return stringsList
}
