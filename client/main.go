package main

import (
	"net"
	"time"
)

func main() {
	remoteAddress, err := net.ResolveUDPAddr("udp4", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp4", nil, remoteAddress)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	tick := time.NewTicker(1 * time.Second)

	for {
		<-tick.C
		conn.Write([]byte("Hello world"))
	}
}