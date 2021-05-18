package main

import (
	"fmt"
)

func main() {
	arr := [10]int{11,12,13,14,15,16,17,18,19,20}
	for i,v := range arr {
		fmt.Println(i,v)	
	}
	fmt.Printf("%T",arr)
}
