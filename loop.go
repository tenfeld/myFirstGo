package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for i < 10 {
		z -= (z*z - x) / (2 * z)

		fmt.Println(z)
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
