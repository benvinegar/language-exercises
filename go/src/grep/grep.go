package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const escapeLightRed = "\u001b[1;31m"
const escapeEnd = "\u001b[0m"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	countLines := flag.Bool("count", false, "Only a count of selected lines is written to standard output.")

	flag.Parse()

	pattern := flag.Arg(0)
	dir := flag.Arg(1)

	data, err := ioutil.ReadFile(dir)
	check(err)

	reader := bufio.NewReader(bytes.NewReader(data))
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var onMatch func(line string)
	var onEnd = func() {}
	if *countLines {
		count := 0
		onMatch = func(line string) {
			count++
		}
		onEnd = func() {
			writer.WriteString(strconv.Itoa(count) + "\n")
		}
	} else {
		patternSub := escapeLightRed + pattern + escapeEnd
		onMatch = func(line string) {
			replaced := strings.ReplaceAll(string(line), pattern, patternSub)
			_, err := writer.WriteString(replaced)
			check(err)
		}
	}

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		check(err)

		if strings.Contains(string(line), pattern) {
			onMatch(string(line))
		}
	}

	onEnd()
}
