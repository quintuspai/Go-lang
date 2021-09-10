//A “callback” is when we pass a func into a func as an argument. 
//pass a func into a func as an argument 



package main

import "fmt"

func main() {
	defer run(hello)
	run(test)
}

func hello() {
	fmt.Println("hello world")	
}

func test() {
	fmt.Println("this is a test")
}

func run(f func()) {
	f()
}
