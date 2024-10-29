package main

import (
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Server Running at localhost:4221")
	l, err := net.Listen("tcp", ":4221")

	if err != nil {
		log.Println("Failed to start server")
		os.Exit(69)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Unable to handle connection")
			os.Exit(60)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("Client connected with addr: %s", conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	log.Println(n)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Printf("Data [%s]:\n", conn.RemoteAddr().String())
	ParsePacket(buffer[:n])
	defer conn.Close()
}

func ParsePacket(packet []byte) {
	log.Printf("Raw Packet: %x", packet)
	version := packet[0]
	data := packet[1 : len(packet)-32]
	hash := packet[len(packet)-33:]
	log.Printf("Version: %d, Data: %s, Hash: %x\n", version, string(data), hash)
}
