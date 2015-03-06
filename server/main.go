package main

import (
	"log"
	"net"
)

func main() {
	address, err := net.ResolveUDPAddr("udp4", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
		}
		if n > 0 {
			log.Printf("Message received from %s: %s", addr.String(), string(buffer[0:n]))
		}
	}
}
