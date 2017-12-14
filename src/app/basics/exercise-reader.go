package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}
type MyAAA string
func (a MyAAA) Error() string {
	return fmt.Sprintf("%s", a)
}

func (m MyReader) Read(b []byte) (int, error) {
	if len(b) <= 0 {
		return len(b), error(MyAAA("it is error!"))
	}

	for index := range b {
		b[index] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}