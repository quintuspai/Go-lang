// use of mutex

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	runs := 100
	counter := 0
	wg.Add(runs)
	for i := 0; i < runs; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
