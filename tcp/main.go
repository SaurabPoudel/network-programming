package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))

	conn.Write([]byte("Message received"))
}

func main() {
	/*
	   net.Conn is a interface that represents a generic stream-oriented network connection
	   this can be used for both TCP and UDP
	   once the connection is established you can use methods like Read, Write and Close to handle the data and connection
	*/
	/*
		net.Listener is a interface that represents a network listener(like a TCP Server) that
		listens for incomming connection on a specific port.
		It provides methods like Accepts to accept the incomming connections and Close to stop
		to the listener
	*/

	// Example

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}
