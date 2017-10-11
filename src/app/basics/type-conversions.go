package main

import (
  "fmt"
  "math"
)

func main() {
  normal()
  shorted()
}

func normal() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}

func shorted() {
  x, y := 3, 4
  f := math.Sqrt(float64(x*x + y*y))
  z := uint(f)
  fmt.Println(x, y, z)
}
