package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func main() {
	absers := []Abser{
		MyFloat(-math.Sqrt2),
		&Vertex{3, 4},
	}

	for _, abser := range absers {
		fmt.Printf("%T\n", abser)
		fmt.Println(abser.Abs())
	}
}
