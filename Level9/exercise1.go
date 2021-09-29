// how to use waitGroups and lunch go subroutines

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo : ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar : ", i)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start...")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
