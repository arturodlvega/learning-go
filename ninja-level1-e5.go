package main

import (
	"fmt"
)

type hola int

var x hola
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x) // CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
