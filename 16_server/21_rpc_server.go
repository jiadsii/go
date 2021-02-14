package main

//import (
//	"net/http"
//	"log"
//	"net"
//	"net/rpc"
//	"time"
//)
//
//type Args struct {
//	N, M int
//}
//
//func (t *Args) Multiply(args *Args, reply *int) error {
//	*reply = args.N * args.M
//	return nil
//}
//
//func main() {
//	calc := new(Args)
//	rpc.Register(calc)
//	rpc.HandleHTTP()
//	listener, e := net.Listen("tcp", "localhost:1234")
//	if e != nil {
//		log.Fatal("Starting RPC-server -listen error:", e)
//	}
//	go http.Serve(listener, nil)
//	time.Sleep(1000e9)
//}