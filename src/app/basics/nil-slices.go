package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if nil == s {
		fmt.Println("nil!")
	}
}
