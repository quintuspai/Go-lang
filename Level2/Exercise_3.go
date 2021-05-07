package main

import "fmt"

const (
	a int = 4
	b = "Hi"	
)


func main() {
	fmt.Println(a,b)
	fmt.Printf("%T\t%T\n",a,b)
}
