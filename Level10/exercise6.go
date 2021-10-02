//comma ok idiom starting with

package main

import "fmt"

func main() {
	c := make(chan int)
	send(c)
	for v := range c {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i + 1
		}
		close(c)
	}()
}
