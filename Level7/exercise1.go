// pointers

package main

import "fmt"

type person struct {
	first string
	middle string
	last string
}

func changeMe(p *person, name string) {
	(*p).middle = name
} 

func main() {
	p1 := person {
		first : "Lionel",
		middle : "test",
		last : "Messi",
	}
	fmt.Println(p1)
	changeMe(&p1, "Andres")
	fmt.Println(p1)
}
