// atomic

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	runs := 100
	var counter uint64
	wg.Add(runs)
	for i := 0; i < runs; i++ {
		go func() {
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
