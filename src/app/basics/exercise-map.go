package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, ok := res[v]
		if ok {
			res[v] += 1
		} else {
			res[v] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
