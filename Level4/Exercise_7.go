package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	z := [][]string{x,y}
	for _,val := range z {
		for _,vals := range val {
			fmt.Println(vals)
		}
		fmt.Println()
	}
	
}