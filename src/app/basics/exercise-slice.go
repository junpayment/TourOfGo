package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := range res {
		temp := make([]uint8, dx)
		for x := range temp {
			temp[x] = uint8((x+y)/2)
		}
		res[y] = temp
	}

	return res
}

func main() {
	pic.Show(Pic)
}



