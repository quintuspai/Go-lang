package main

import "fmt"

func main() {
	var favSport string = "football"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "football":
		fmt.Println("go to pitch!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
