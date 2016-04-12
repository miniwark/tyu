package main

import "strings"

// Convert and concatenate a []int8 slice to a string
// inspired by perterGo: https://groups.google.com/forum/#!topic/golang-nuts/Jel8Bb-YwX8
func int8SliceToString(char []int8) string {
	s := make([]string, len(char))
	for i := range s {
		s[i] = string(char[i])
	}
	return strings.Join(s, "")
}

// TODO maybe change this to convert exclusively [65]int8 array ?
