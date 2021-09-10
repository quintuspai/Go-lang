//Closure is when we have “enclosed” the scope of a variable in some code block.
//create a func which “encloses” the scope of a variable:

package main

import "fmt"

func main() {
	f := fun()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := fun()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	
}

func fun() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}