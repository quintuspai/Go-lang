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
	
	m := map[string]person{
		p1.first : p1,
		p2.first : p2,
	}
	
	for _,v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i,v1 := range v.icecream {
			fmt.Println(i,v1)
		}
		fmt.Println() 