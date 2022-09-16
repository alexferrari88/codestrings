package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
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

	output := flag.String("output", "csv", "Output type: json, csv")
	flag.Parse()
	filesArg := flag.Args()
	if len(filesArg) == 0 {
		panic("No files specified")
	}

	type Result struct {
		File   string   `json:"file"`
		Data   []string `json:"data"`
		Output string   `json:"-"` // This field will not be marshalled
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(filesArg))

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
			results <- Result{
				File:   file,
				Data:   ExtractStrings(string(raw), []string{"\"", "'", "`"}),
				Output: *output,
			}
		}(file)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		switch result.Output {
		case "json":
			j, err := json.Marshal(result)
			if err != nil {
				continue
			}
			fmt.Println(string(j))
		case "csv":
			fmt.Printf("%s,%s\n", result.File, strings.Join(result.Data, ","))
		default:
			panic("Unknown output format")
		}
	}
}
