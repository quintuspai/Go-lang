
package main

import "fmt"

func main() {
	c := make(chan int)
	send(c)

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

}

func send(c chan<- int) {

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j + 1
			}
		}()
	}
}
