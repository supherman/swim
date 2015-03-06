package main

import (
	"flag"

	"github.com/supherman/swim/server"
)

var port = flag.String("port", "0", "3000")
var peer = flag.String("peer", "", "127.0.0.1:4000")

func main() {
	flag.Parse()
	server, err := server.New(*port, *peer)
	if err != nil {
		panic(err)
	}

	server.Run()
}
