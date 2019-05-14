package main

import "net/rpc"
import "net"
import "errors"
import "fmt"
import "net/http"
import "log"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	fmt.Println("Here")
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// rpc.Accept(l)
	fmt.Println("Here")
	err := http.Serve(l, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}

}