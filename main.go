package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var filters map[string]string = {"": ""}

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
