package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Vertex:
		fmt.Println("hoge")

	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(Vertex{2,3})
}
