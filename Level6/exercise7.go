//Assign a func to a variable, then call that func


package main

import "fmt"

func main() {
	display := func (n int) {
		fmt.Println(n)
	}
	
	display(2)
}
