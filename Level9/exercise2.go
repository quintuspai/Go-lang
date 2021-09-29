//method sets


package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am ")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Lionel",
		last:  "Messi",
	}
	saySomething(&p1)

}
