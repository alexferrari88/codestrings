package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	cs "github.com/alexferrari88/codestrings"
)

func main() {

	delimitersFlag := flag.String("delimiters", "\",',`", "delimiters to use for string extraction (comma separated and escaped)")
	output := flag.String("output", "csv", "output type: json, csv")
	flag.Parse()
	filesArg := flag.Args()
	if len(filesArg) == 0 {
		fmt.Println("No files provided")
		os.Exit(1)
	}
	delimiters := strings.Split(*delimitersFlag, ",")

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
				fmt.Println(err.Error())
				os.Exit(1)
			}
			results <- Result{
				File:   file,
				Data:   cs.ExtractStrings(string(raw), delimiters),
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
			for i := range result.Data {
				result.Data[i] = strings.Replace(result.Data[i], "\"", "\"\"", -1)
				result.Data[i] = fmt.Sprintf("\"%s\"", result.Data[i])
			}
			csvData := strings.Join(result.Data, ",")
			fmt.Printf("%s,%s\n", result.File, csvData)
		default:
			continue
		}
	}
}
