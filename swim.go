package main

import (
	"flag"
	"time"

	"github.com/supherman/swim/swim"
)

var port = flag.String("port", "", "3000")
var initial = flag.String("initial-peer", "", "4000")

func main() {
	flag.Parse()
	server, err := swim.NewServer(*port, *initial, 10*time.Second)
	if err != nil {
		panic(err)
	}

	server.Run()
}
