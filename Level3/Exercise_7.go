package main

import "fmt"

func main() {
	x := ""
	if x == "Thomas" {
		fmt.Println(x)
	} else if x == "Pingu" {
		fmt.Println("Noot Noot")
	} else {
		fmt.Println(":(")
	}
}