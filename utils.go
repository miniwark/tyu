package main

import (
	"io/ioutil"
	"strings"
)

// Convert and concatenate a []int8 slice to a string
// inspired by perterGo: https://groups.google.com/forum/#!topic/golang-nuts/Jel8Bb-YwX8
func int8SliceToString(char []int8) string {
	s := make([]string, len(char))
	for i := range s {
		s[i] = string(char[i])
	}
	return strings.Join(s, "")
}

// TODO maybe change this to convert exclusively to [65]int8 array ?

// Read a text file and return the content without the EOF carriage return
// this utility may be used with files from `/sys` or `/proc` file systems
func readAndTrimFile(path string) string {
	ret := ""
	data, err := ioutil.ReadFile(path)
	if err == nil {
		ret = strings.TrimRight(string(data), "\n") // Trimright remove the EOF carriage return
	}
	return ret
}

//TODO manage multiple line files ?
