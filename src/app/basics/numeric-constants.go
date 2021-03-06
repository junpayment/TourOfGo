package main

import "fmt"

const (
  Big = 1 << 100
  Small = 1 >> 99
)

func needInt(x int) int {
  fmt.Printf("%T\n", x)
  return x * 10 + 1
}

func needFloat(x float64) float64 {
  fmt.Printf("%T\n", x)
  return x * 0.1
}

func main() {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}
