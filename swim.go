package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.String("port", "0", "3000")
var peer = flag.String("peer", "", "127.0.0.1:4000")

func main() {
	flag.Parse()
	addresString := fmt.Sprintf("127.0.0.1:%s", *port)
	address, err := net.ResolveUDPAddr("udp4", addresString)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", address)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening at %s", address.String())

	defer conn.Close()

	if *peer != "" {
		peerAddress, err := net.ResolveUDPAddr("udp4", *peer)
		if err != nil {
			panic(err)
		}

		peerConn, err := net.DialUDP("udp4", nil, peerAddress)
		if err == nil {
			peerConn.Write([]byte(address.String()))
		} else {
			panic(err)
		}
	}

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
		}
		if n > 0 {
			clientAddress, err := net.ResolveUDPAddr("udp4", string(buffer))
			if err == nil {
				clientConn, err := net.DialUDP("udp4", nil, clientAddress)
				if err == nil {
					clientConn.Write([]byte(fmt.Sprintf("ACK: %s", string(buffer))))
				}
			}
			log.Println(string(buffer))
		}
	}
}
