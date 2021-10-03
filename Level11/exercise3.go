// custom error

package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("the error ; %v", ce.info)
}

func main() {
	c := customErr{
		info: "need more tea",
	}
	foo(c)
	fmt.Println("")
}

func foo(e error) {
	fmt.Println("foo ran-", e)
}
