package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sync"
)

func populateMap(m map[string]struct{}, s []string) {
	for _, v := range s {
		m[v] = struct{}{}
	}
}

func ExtractStrings(source string, stringDelimiters []string) []string {
	var strings []string
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
				strings = append(strings, currentString)
				started = false
				currentString = ""
				continue
			} else {
				started = true
				continue
			}
		}
		if started {
			currentString += string(c)
		}
	}
	return strings
}

func main() {

	flag.Parse()
	filesArg := flag.Args()
	if len(filesArg) == 0 {
		panic("No files specified")
	}

	var wg sync.WaitGroup

	for _, file := range filesArg {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Error processing file %s: %v", file, err)
				}
			}()

			raw, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%q\n", ExtractStrings(string(raw), []string{"\"", "'", "`"}))
		}(file)
	}

	wg.Wait()
}
