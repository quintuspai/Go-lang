package main

import (
	"fmt"
)

func main() {
	arr := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])
}