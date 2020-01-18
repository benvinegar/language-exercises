package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pattern := os.Args[1]
	dir := os.Args[2]

	patternSub := "\u001b[1;31m" + pattern + "\u001b[0m"

	data, err := ioutil.ReadFile(dir)
	check(err)

	reader := bufio.NewReader(bytes.NewReader(data))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		check(err)

		if strings.Contains(string(line), pattern) {
			replaced := strings.ReplaceAll(string(line), pattern, patternSub)
			fmt.Println(replaced)
		}
	}
}
