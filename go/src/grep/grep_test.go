package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {

	data, _ := ioutil.ReadFile("fixtures/loremipsum.txt")
	reader := bufio.NewReader(bytes.NewReader(data))

	buf := bytes.NewBuffer(make([]byte, 0)) // loremipsum.txt is 4011 chars
	writer := bufio.NewWriter(buf)

	Grep("accumsan", reader, writer, GrepOpts{})

	writer.Flush()

	result := buf.String()
	expected := "Phasellus eros. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Donec nibh. Integer \u001b[1;31maccumsan\u001b[0m, purus semper scelerisque vestibulum, elit wisi rhoncus arcu, et pulvinar augue justo ut dui. Duis consectetuer, sem at hendrerit commodo, dolor ligula tristique mi, vitae venenatis sapien sem sed urna. Mauris vehicula purus sed wisi. Vestibulum odio. Donec arcu est, dignissim ac, convallis sit amet, suscipit quis, eros. Vestibulum dictum elit sit amet leo. Vivamus dignissim. Aenean quam. Phasellus semper diam. Quisque bibendum ullamcorper neque. "
	assert.Equal(t, expected, result)
}
