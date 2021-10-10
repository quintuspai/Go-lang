//server
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Number struct {
	A   int
	B   int
	Ops int
}

type API int

func (a *API) GetOperationList(payload string, reply *[]string) error {
	ops := []string{"Addtion", "Substraction", "Multiplication", "Division", "Modulo"}
	*reply = ops
	return nil
}

func (a *API) Perform(payload Number, reply *int) error {
	var x int
	switch payload.Ops {
	case 1:
		x = payload.A + payload.B
	case 2:
		x = payload.A - payload.B
	case 3:
		x = payload.A * payload.B
	case 4:
		x = payload.A / payload.B
	case 5:
		x = payload.A % payload.B
	}
	*reply = int(x)
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering rpc", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listener error:", err)
	}

	log.Printf("On Port : %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}
