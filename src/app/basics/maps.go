package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// 連想配列的な？
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

