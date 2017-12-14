package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (count int, err error) {
	count, err = rot.r.Read(b)
	if io.EOF == err {
		return
	}
	for index, org := range b {
		// 大文字,小文字で場合分けしてさらにM|mで場合分け
		if org >= 0x41 && org <= 0x5a {
			if org <= 0x4D {
				b[index] = org + 13
			} else {
				b[index] = org - 13
			}
		} else if org >= 0x61 && org <= 0x7a {
			if org <= 0x6D {
				b[index] = org + 13
			} else {
				b[index] = org - 13
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

