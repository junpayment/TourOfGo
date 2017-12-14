package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
  //return fmt.Sprintf("cannot Sqrt negative number: %v", e) // 無限ループ
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := float64(1)
	for n := 0; n < 10; n++ {
		z = z - ((math.Pow(z, 2.0) - x) / (2 * z))
	}
	return z, nil
}

type MyInt int
func (i MyInt) Error() string {
	return fmt.Sprintf("aaaaaaa %v", int(i))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	var ie error
	var v MyInt = 1
	ie = v
	fmt.Println(ie)
}
