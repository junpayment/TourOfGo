package main

import (
	"fmt"
	"math"
)

// class Vertex {
//  float64 X, Y;
//  function Abs() float64 {
//  	return math.Sqrt(X*X + Y*Y)
//  }
//}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
