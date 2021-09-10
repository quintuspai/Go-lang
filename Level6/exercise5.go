// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle


package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64{
	return math.Pi * 2 * c.radius
}

func (s square) area() float64{
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}


func main() {
	circ := circle{
		radius : 12.234,
	}
	
	squ := square {
		length : 4.0,
	}
	info(circ)
	info(squ)
}
