package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrepBasic(t *testing.T) {
	data, _ := ioutil.ReadFile("fixtures/loremipsum.txt")
	reader := bufio.NewReader(bytes.NewReader(data))

	buf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf)

	// no matches
	Grep("nomatches", reader, writer, GrepOpts{})
	writer.Flush()

	assert.Equal(t, "", buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// a single match
	Grep("accumsan", reader, writer, GrepOpts{})
	writer.Flush()

	expected := "Phasellus eros. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Donec nibh. Integer \u001b[1;31maccumsan\u001b[0m, purus semper scelerisque vestibulum, elit wisi rhoncus arcu, et pulvinar augue justo ut dui. Duis consectetuer, sem at hendrerit commodo, dolor ligula tristique mi, vitae venenatis sapien sem sed urna. Mauris vehicula purus sed wisi. Vestibulum odio. Donec arcu est, dignissim ac, convallis sit amet, suscipit quis, eros. Vestibulum dictum elit sit amet leo. Vivamus dignissim. Aenean quam. Phasellus semper diam. Quisque bibendum ullamcorper neque. \n"
	assert.Equal(t, expected, buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// multiple matches
	Grep("ultrices", reader, writer, GrepOpts{})
	writer.Flush()

	expected = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vestibulum ante ipsum primis in faucibus orci luctus et \u001b[1;31multrices\u001b[0m posuere cubilia Curae; Mauris ornare, felis eget lacinia congue, ante lorem condimentum lacus, non aliquet leo erat sed sem. Proin magna. Nullam est. Phasellus non risus semper velit blandit mollis. Suspendisse neque ante, facilisis ac, cursus vel, fermentum a, sapien. In hac habitasse platea dictumst. Curabitur augue pede, placerat vel, lacinia vitae, euismod a, ante. In hac habitasse platea dictumst. Integer diam nibh, varius ut, scelerisque sit amet, sagittis quis, tortor. Nullam eget magna. Praesent pede nibh, molestie nec, malesuada in, venenatis semper, diam. Nunc luctus eleifend lorem. Aliquam ut magna sed dui vestibulum dignissim. Maecenas vitae elit. Aenean nulla augue, pulvinar id, malesuada semper, dignissim sit amet, velit. Duis et quam sed justo pretium elementum. Curabitur risus. Ut hendrerit facilisis felis. Fusce molestie volutpat erat. Aenean lobortis magna et orci. Aenean nec est sit amet ligula auctor dictum. Maecenas pretium lectus ut magna. Integer ornare sollicitudin metus. \n" +
		"Sed nunc. Nullam vitae mauris sed libero laoreet posuere. Morbi rhoncus. Duis id enim nec sapien fringilla volutpat. Morbi hendrerit, nulla sit amet consectetuer scelerisque, lacus leo sodales est, vel tempor sem leo nec lacus. Nam ligula. Sed non orci. Vestibulum ante ipsum primis in faucibus orci luctus et \u001b[1;31multrices\u001b[0m posuere cubilia Curae; Sed id metus. Praesent augue. Fusce a mi. Nulla nonummy. \n" +
		"Donec ac mauris id metus faucibus aliquet. Duis semper lorem et diam. Donec eget velit. In scelerisque. Cras ut tellus vitae dui vestibulum rhoncus. Sed vel lectus eu est aliquam scelerisque. Donec et sapien et arcu scelerisque vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et \u001b[1;31multrices\u001b[0m posuere cubilia Curae; Nulla facilisi. Duis nulla. Pellentesque magna odio, auctor in, malesuada vel, hendrerit ut, neque. Aenean iaculis luctus orci. Aliquam vel neque ut diam congue convallis. Praesent \u001b[1;31multrices\u001b[0m, urna nec commodo pharetra, enim pede euismod sem, vitae euismod libero sapien non lacus. Vivamus velit. Suspendisse quis mauris. Phasellus rutrum dolor vitae nisl. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nunc aliquet, sapien at ullamcorper facilisis, quam odio gravida mi, vitae molestie lectus massa eget neque. Phasellus hendrerit condimentum mauris. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. \n"
	assert.Equal(t, expected, buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// case insensitive
	Grep("vestibulum", reader, writer, GrepOpts{caseInsensitive: true})
	writer.Flush()

	expected = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. \u001b[1;31mVestibulum\u001b[0m ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris ornare, felis eget lacinia congue, ante lorem condimentum lacus, non aliquet leo erat sed sem. Proin magna. Nullam est. Phasellus non risus semper velit blandit mollis. Suspendisse neque ante, facilisis ac, cursus vel, fermentum a, sapien. In hac habitasse platea dictumst. Curabitur augue pede, placerat vel, lacinia vitae, euismod a, ante. In hac habitasse platea dictumst. Integer diam nibh, varius ut, scelerisque sit amet, sagittis quis, tortor. Nullam eget magna. Praesent pede nibh, molestie nec, malesuada in, venenatis semper, diam. Nunc luctus eleifend lorem. Aliquam ut magna sed dui \u001b[1;31mvestibulum\u001b[0m dignissim. Maecenas vitae elit. Aenean nulla augue, pulvinar id, malesuada semper, dignissim sit amet, velit. Duis et quam sed justo pretium elementum. Curabitur risus. Ut hendrerit facilisis felis. Fusce molestie volutpat erat. Aenean lobortis magna et orci. Aenean nec est sit amet ligula auctor dictum. Maecenas pretium lectus ut magna. Integer ornare sollicitudin metus. \n" +
		"Sed nunc. Nullam vitae mauris sed libero laoreet posuere. Morbi rhoncus. Duis id enim nec sapien fringilla volutpat. Morbi hendrerit, nulla sit amet consectetuer scelerisque, lacus leo sodales est, vel tempor sem leo nec lacus. Nam ligula. Sed non orci. \u001b[1;31mVestibulum\u001b[0m ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Sed id metus. Praesent augue. Fusce a mi. Nulla nonummy. \n" +
		"Donec ac mauris id metus faucibus aliquet. Duis semper lorem et diam. Donec eget velit. In scelerisque. Cras ut tellus vitae dui \u001b[1;31mvestibulum\u001b[0m rhoncus. Sed vel lectus eu est aliquam scelerisque. Donec et sapien et arcu scelerisque \u001b[1;31mvestibulum\u001b[0m. \u001b[1;31mVestibulum\u001b[0m ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla facilisi. Duis nulla. Pellentesque magna odio, auctor in, malesuada vel, hendrerit ut, neque. Aenean iaculis luctus orci. Aliquam vel neque ut diam congue convallis. Praesent ultrices, urna nec commodo pharetra, enim pede euismod sem, vitae euismod libero sapien non lacus. Vivamus velit. Suspendisse quis mauris. Phasellus rutrum dolor vitae nisl. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nunc aliquet, sapien at ullamcorper facilisis, quam odio gravida mi, vitae molestie lectus massa eget neque. Phasellus hendrerit condimentum mauris. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. \n" +
		"Phasellus eros. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Donec nibh. Integer accumsan, purus semper scelerisque \u001b[1;31mvestibulum\u001b[0m, elit wisi rhoncus arcu, et pulvinar augue justo ut dui. Duis consectetuer, sem at hendrerit commodo, dolor ligula tristique mi, vitae venenatis sapien sem sed urna. Mauris vehicula purus sed wisi. \u001b[1;31mVestibulum\u001b[0m odio. Donec arcu est, dignissim ac, convallis sit amet, suscipit quis, eros. \u001b[1;31mVestibulum\u001b[0m dictum elit sit amet leo. Vivamus dignissim. Aenean quam. Phasellus semper diam. Quisque bibendum ullamcorper neque. \n" +
		"Nunc tempor nibh ac eros. Proin magna nunc, laoreet vitae, scelerisque in, tincidunt a, arcu. Fusce semper lorem. Nulla non est vel elit condimentum sagittis. Pellentesque fermentum libero semper augue. Morbi nibh. Praesent tellus. Phasellus ac massa. Duis nec augue. Aliquam semper dignissim nisl. Integer sed nulla vel nisl blandit venenatis. Nam velit. Aliquam nulla. Integer ac magna vitae libero imperdiet \u001b[1;31mvestibulum\u001b[0m. Morbi vel urna vitae justo sagittis laoreet. \n"
	assert.Equal(t, expected, buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// max count
	Grep("vestibulum", reader, writer, GrepOpts{caseInsensitive: true, maxCount: 1, maxCountEnabled: true})
	writer.Flush()

	expected = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. \u001b[1;31mVestibulum\u001b[0m ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris ornare, felis eget lacinia congue, ante lorem condimentum lacus, non aliquet leo erat sed sem. Proin magna. Nullam est. Phasellus non risus semper velit blandit mollis. Suspendisse neque ante, facilisis ac, cursus vel, fermentum a, sapien. In hac habitasse platea dictumst. Curabitur augue pede, placerat vel, lacinia vitae, euismod a, ante. In hac habitasse platea dictumst. Integer diam nibh, varius ut, scelerisque sit amet, sagittis quis, tortor. Nullam eget magna. Praesent pede nibh, molestie nec, malesuada in, venenatis semper, diam. Nunc luctus eleifend lorem. Aliquam ut magna sed dui \u001b[1;31mvestibulum\u001b[0m dignissim. Maecenas vitae elit. Aenean nulla augue, pulvinar id, malesuada semper, dignissim sit amet, velit. Duis et quam sed justo pretium elementum. Curabitur risus. Ut hendrerit facilisis felis. Fusce molestie volutpat erat. Aenean lobortis magna et orci. Aenean nec est sit amet ligula auctor dictum. Maecenas pretium lectus ut magna. Integer ornare sollicitudin metus. \n"
	assert.Equal(t, expected, buf.String())
}

func TestGrepCount(t *testing.T) {
	data, _ := ioutil.ReadFile("fixtures/loremipsum.txt")
	reader := bufio.NewReader(bytes.NewReader(data))

	buf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf)

	// no matches
	Grep("nomatches", reader, writer, GrepOpts{countLines: true})
	writer.Flush()

	assert.Equal(t, "0\n", buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// a single match
	Grep("accumsan", reader, writer, GrepOpts{countLines: true})
	writer.Flush()

	assert.Equal(t, "1\n", buf.String())

	reader.Reset(bytes.NewReader(data))
	buf.Reset()

	// multiple matches
	Grep("ultrices", reader, writer, GrepOpts{countLines: true})
	writer.Flush()

	assert.Equal(t, "3\n", buf.String()) // 3 lines (but 4 matches total; --count outputs lines matched)
}
