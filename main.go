package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := flag.String("file", "", "File path")
	flag.Parse()

	raw, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err.Error())
	}
	var strings []string
	var started bool
	var currentString string
	for i := range raw {
		c := raw[i]
		if c == '`' || c == '"' || c == '\'' {
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
	fmt.Println(strings)
}
