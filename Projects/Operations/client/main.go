package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type Number struct {
	A   int
	B   int
	Ops int
}

func main() {
	var reply int

	var ops []string

	scanner := bufio.NewScanner(os.Stdin)

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client.Call("API.GetOperationList", "", &ops)

	for {
		fmt.Println("select choice by entering number. Enter 0 to exit")
		for i, val := range ops {
			fmt.Println(i+1, ":", val)
		}
		scanner.Scan()
		switch scanner.Text() {
		case "0":
			os.Exit(0)
		case "1":
			fmt.Println("You Selected Addition.")
			fmt.Println("Enter Number 1: ")
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter Number 2: ")
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			a := Number{
				x,
				y,
				1,
			}
			client.Call("API.Perform", a, &reply)
			fmt.Println(x, "+", y, "=", reply)
		case "2":
			fmt.Println("You Selected Substraction.")
			fmt.Println("Enter Number 1: ")
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter Number 2: ")
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			a := Number{
				x,
				y,
				2,
			}
			client.Call("API.Perform", a, &reply)
			fmt.Println(x, "-", y, "=", reply)
		case "3":
			fmt.Println("You Selected Multiplication.")
			fmt.Println("Enter Number 1: ")
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter Number 2: ")
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			a := Number{
				x,
				y,
				3,
			}
			client.Call("API.Perform", a, &reply)
			fmt.Println(x, "*", y, "=", reply)
		case "4":
			fmt.Println("You Selected Division.")
			fmt.Println("Enter Number 1: ")
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter Number 2: ")
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			a := Number{
				x,
				y,
				4,
			}
			client.Call("API.Perform", a, &reply)
			fmt.Println(x, "/", y, "=", reply)
		case "5":
			fmt.Println("You Selected Modulo.")
			fmt.Println("Enter Number 1: ")
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter Number 2: ")
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			a := Number{
				x,
				y,
				5,
			}
			client.Call("API.Perform", a, &reply)
			fmt.Println(x, "+", y, "%", reply)
		}
	}
}
