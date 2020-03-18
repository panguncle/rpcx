package main

import (
	"flag"

	"github.com/smallnest/rpcx/server"
	"github.com/smallnestg/rpcx/example"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	//s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}
