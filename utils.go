package main

import (
	"strconv"
	"strings"
)

// convert and concatenate a []int8 array to a string
// inspired by perterGo: https://groups.google.com/forum/#!topic/golang-nuts/Jel8Bb-YwX8
func int8ArrayToString(char []int8) string {

	s := make([]string, len(char))

	for i := range s {
		s[1] = strconv.QuoteRune(rune(char[i])) // convert each int8 char to a rune
	}
	return strings.Join(s, "")
}
