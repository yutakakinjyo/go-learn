package main

import "code.google.com/p/go-tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, k := range strings.Fields(s) {
		m[k] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
