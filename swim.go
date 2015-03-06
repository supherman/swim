package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.String("port", "0", "3000")
var peer = flag.String("peer", "", "127.0.0.1:4000")

type Server struct {
	Address *net.UDPAddr
	Peers   []*net.UDPAddr
}

func NewServer(port, peerPort string) (*Server, error) {
	address, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, err
	}

	peers := []*net.UDPAddr{}

	peerAddress, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", peerPort))
	if err == nil {
		peers = append(peers, peerAddress)
	} else {
		log.Println(err)
	}

	return &Server{
		Address: address,
		Peers:   peers,
	}, nil
}

func (s *Server) Run() {
	if len(s.Peers) > 0 {
		s.Ping(s.Peers[0])
	}

	conn, err := net.ListenUDP("udp4", s.Address)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening at %s", s.Address.String())

	defer conn.Close()

	for {
		message := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(message)
		if err != nil {
			log.Println(err)
		}

		if n > 0 {
			clientAddress, err := net.ResolveUDPAddr("udp4", string(message[0:n]))
			if err == nil {
				s.Ack(clientAddress)
			}
			log.Println(string(message[0:n]))
		}
	}
}

func (s *Server) Ping(addr *net.UDPAddr) error {
	clientConn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return err
	}
	log.Println("Pinging: ", s.Address.String())
	clientConn.Write([]byte(s.Address.String()))
	return nil
}

func (s *Server) Ack(addr *net.UDPAddr) error {
	clientConn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return err
	}
	log.Println("Responding back to: ", s.Address.String())
	clientConn.Write([]byte(fmt.Sprintf("ACK: %s", s.Address.String())))
	return nil
}

func main() {
	flag.Parse()
	server, err := NewServer(*port, *peer)
	if err != nil {
		panic(err)
	}

	server.Run()
}
