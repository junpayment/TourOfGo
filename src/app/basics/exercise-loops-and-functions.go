package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  for n := 0; n < 10; n++ {
    z = z - ((math.Pow(z, 2.0) - x) / (2 * z))
  }
  return z
}

func main() {
  fmt.Println(math.Sqrt(2))
  fmt.Println(Sqrt(2))
}
