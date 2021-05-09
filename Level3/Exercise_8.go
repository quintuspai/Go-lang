package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Won't print")
	case true:
		fmt.Println("Hello")
	}

}