package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// grupperar en mängd variabler vilka kan anropas med sammlingsnamet
type Struct struct {
	M, N int64
}

// Reads data from the connection
func handleConnection(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	pointer := &struct{}
	decoder.Decode(pointer)
	fmt.Println("Received message : %+v", pointer);
}

func main() {
	fmt.Println("Start listening after clients...");

	ln, err := net.Listen("tcp", ":8080")
	if er != nil {
		fmt.Println("tcp or port problem");
	}
	for {
		conn, err := ln.Accept() // blockerar till den får kontakt eller ett felmeddelande visas
		if err != nil {
			fmt.Println("No connection with the client could be made");
			continue
		}
		go handleConnection(conn)
	}
}
