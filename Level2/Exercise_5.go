package main

import "fmt"

//Using iota, create 4 constants for the NEXT 4 years. 
const (
	a = 2021 + iota
	b  
	c 
	d 
	 
)

func main() {
	fmt.Println(a,b,c,d)
}
