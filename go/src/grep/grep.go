package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

const escapeLightRed = "\u001b[1;31m"
const escapeEnd = "\u001b[0m"
const programName = "grep"

// GrepOpts holds options passed to the Grep function
type GrepOpts struct {
	countLines      bool
	caseInsensitive bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	countLines := flag.Bool("count", false, "Only a count of selected lines is written to standard output.")
	caseInsensitive := flag.Bool("ignore-case", false, "Perform case insensitive matching.  By default, grep is case sensitive.")

	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	opts := GrepOpts{*countLines, *caseInsensitive}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", programName, err)
	}

	reader := bufio.NewReader(bytes.NewReader(data))
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	Grep(pattern, reader, writer, opts)
}

// Grep searches the file located at `path` for matches of `pattern`
func Grep(pattern string, reader *bufio.Reader, writer *bufio.Writer, opts GrepOpts) {

	var patternRe *regexp.Regexp
	if opts.caseInsensitive {
		patternRe, _ = regexp.Compile("((?i)" + pattern + ")")
	} else {
		patternRe, _ = regexp.Compile("(" + pattern + ")")
	}

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
		patternSub := escapeLightRed + "$1" + escapeEnd
		onMatch = func(line string) {
			replaced := patternRe.ReplaceAllString(string(line), patternSub)
			_, err := writer.WriteString(replaced + "\n")
			check(err)
		}
	}

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		check(err)

		if patternRe.MatchString(string(line)) {
			onMatch(string(line))
		}
	}

	onEnd()
}
