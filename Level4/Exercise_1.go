package main

import (
	"fmt"
)

func main() {
	arr := [5]int{41,42,43,44,45}
	for i, val := range arr {
		fmt.Println(i,val)
	}
	fmt.Printf("%T",arr)
}