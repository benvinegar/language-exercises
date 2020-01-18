package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const escapeLightRed = "\u001b[1;31m"
const escapeEnd = "\u001b[0m"
const programName = "grep"

type grepOpts struct {
	countLines bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	countLines := flag.Bool("count", false, "Only a count of selected lines is written to standard output.")

	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	opts := grepOpts{*countLines}
	Grep(pattern, path, opts)
}

// Grep searches the file located at `path` for matches of `pattern`
func Grep(pattern string, path string, opts grepOpts) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", programName, err)
	}

	reader := bufio.NewReader(bytes.NewReader(data))
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var onMatch func(line string)
	var onEnd = func() {}
	if opts.countLines {
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
