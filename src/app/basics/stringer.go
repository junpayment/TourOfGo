package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Zephod Beeblebrox", 9001}
	fmt.Println(a, b)

	// see below
	//func Println(a ...interface{}) (n int, err error) {
	//	return Fprintln(os.Stdout, a...)
	//}
}

