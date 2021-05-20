package main

import (
	"fmt"
)

func main() {
	v := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "white",
	}
	fmt.Println(v)
}
