package main

import (
	"fmt"
)

type booshi int

var x booshi

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
