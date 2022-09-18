package codestrings

import (
	"strings"
	"unicode/utf8"
)

type Scanner struct {
	delimitersMap map[string]struct{}
	source        string
	currentRune   rune
	currentPos    int
	nextPos       int
}

func NewScanner(source string, delimitersMap map[string]struct{}) *Scanner {
	var scan = Scanner{delimitersMap, source, -1, 0, 0}
	scan.next()
	return &scan
}

func (scan *Scanner) peek(w int) string {
	if scan.nextPos+w >= len(scan.source) {
		return ""
	}
	return scan.source[scan.nextPos : scan.nextPos+w]
}

func (scan *Scanner) next() {
	if scan.currentPos >= len(scan.source)-1 {
		scan.nextPos = len(scan.source)
		scan.currentRune = -1
		return
	}
	scan.currentPos = scan.nextPos
	r, w := rune(scan.source[scan.nextPos]), 1
	if r >= utf8.RuneSelf {
		r, w = utf8.DecodeRuneInString(scan.source[scan.nextPos:])
	}
	scan.nextPos += w
	scan.currentRune = r
}

func (scan *Scanner) skipSpaces() {
	for scan.currentRune == ' ' || scan.currentRune == '\t' || scan.currentRune == '\n' || scan.currentRune == '\r' {
		scan.next()
	}
}

func (scan *Scanner) scanString(delimiter string) string {
	// move to the next rune until you find the delimiter
	// if you find an escaped delimiter, include the backslash and the delimiter in the result
	// then move to the next rune
	scan.next()
	var startPos = scan.currentPos
	for string(scan.currentRune) != delimiter {
		if scan.currentRune == '\\' && string(scan.peek(len(delimiter))) == delimiter {
			scan.next()
			for i := 0; i < len(delimiter); i++ {
				scan.next()
			}
		}
		scan.next()
	}
	var endPos = scan.currentPos
	scan.next()
	return scan.source[startPos:endPos]
}

func (scan *Scanner) resetTo(pos int) {
	scan.currentPos = pos
	scan.nextPos = pos
	scan.next()
}

func (scan *Scanner) nextString() string {
	// move along the string until you find a delimiter
	// then invoke scanString with the delimiter
	scan.skipSpaces()
	if scan.currentRune < 0 {
		return ""
	}
	for scan.currentRune > 0 {
		if _, ok := scan.delimitersMap[string(scan.currentRune)]; ok {
			return scan.scanString(string(scan.currentRune))
		}
		scan.next()
	}
	return ""
}

func ExtractStrings(source string, stringDelimiters []string) []string {
	var stringsList []string
	if source == "" {
		return stringsList
	}
	if len(stringDelimiters) == 0 {
		stringDelimiters = []string{"\""}
	}

	var delimitersMap = make(map[string]struct{})
	populateMap(delimitersMap, stringDelimiters)
	scan := NewScanner(source, delimitersMap)
	for {
		var str = scan.nextString()
		if str != "" && strings.TrimSpace(str) != "" {
			stringsList = append(stringsList, str)
		}
		if scan.currentRune < 0 {
			break
		}
	}
	return stringsList
}

func populateMap(m map[string]struct{}, s []string) {
	for _, v := range s {
		m[v] = struct{}{}
	}
}
