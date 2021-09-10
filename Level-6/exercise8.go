//function returing a function


package main

import "fmt"

func main() {
	m := nestedfunc()
	fmt.Println(m(2,2))
}

func nestedfunc() func (int, int) int{
	return func (x , y int) int {
		return x * y
	}
}