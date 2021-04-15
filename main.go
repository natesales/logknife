package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
    regexIPv4 = `/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/`
)

var filters map[string]string = {"ipv4": []string{regexIPv4}}

func getIp


func replace(replacement string, pattern string) string {
     return regexp.MustCompile(pattern).ReplaceAllString(replacement)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		fmt.Println(input)
	}
}
