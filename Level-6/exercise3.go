//use of defer


package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	defer foo(nums...)
	fmt.Println("main exit")
}

func foo(a ...int) {
	sum := 0
	for i, _ := range a {
		sum += a[i]
	}
	fmt.Println(sum)
}


