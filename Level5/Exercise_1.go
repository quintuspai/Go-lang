package main

import "fmt"


type person struct{
	first string
	last string
	icecream []string
}

func main() {
	p1 := person {
		first: "James",
		last: "Bond",
		icecream: []string{"Mango","Coconut","Strawberry"}, 
	}
	p2 := person {
		first: "Harley",
		last: "Quinn",
		icecream: []string{"Chocolate","Vanilla","Martini"}, 
	}
	
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.icecream{
		fmt.Println(v)
	}
	
	fmt.Println()
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for _, v := range p2.icecream{
		fmt.Println(v)
	}
}
