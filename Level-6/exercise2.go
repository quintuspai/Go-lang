//create a func with the identifier foo that 
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in
//create a func with the identifier bar that 
//takes in a parameter of type []int
//returns the sum of all values of type int passed in


package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	fmt.Println(foo(nums...))
}

func foo(a ...int) int {
	sum := 0
	for i, _ := range a {
		sum += a[i]
	}
	return sum
}


