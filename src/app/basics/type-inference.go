package main

import (
  "fmt"
)

func main() {
  v := 42
  f := 3.1
  g := 0.8 + 0.5i
  fmt.Printf("%T %T %T", v, f, g)
}
