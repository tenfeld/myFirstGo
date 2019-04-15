package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	result := [dx][dy]uint8{}

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			result[dx][dy] = 0
		}
	}
}

func main() {
	pic.Show(Pic)
}
