package main

import (
	"fmt"
)

type ninja int

var x ninja

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
}