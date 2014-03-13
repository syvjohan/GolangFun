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
	pointer := &Struct{}
	decoder.Decode(pointer)
	fmt.Println("Received message : %+v", pointer)

	sendBackMsg(conn)
}

func sendBackMsg(conn) {
	encoder := gob.NewEncoder(conn) // NewEncoder returns a new encoder that will transmit on the io.Writer.
	pointer := &Struct{1, 2}
	encoder.Encode(pointer)
	fmt.Println("done")
}

func main() {
	fmt.Println("Start listening after clients...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("tcp or port problem")
	}
	for {
		conn, err := ln.Accept() // blockerar till den får kontakt eller ett felmeddelande visas
		if err != nil {
			fmt.Println("No connection with the client could be made")
			continue
		}
		go handleConnection(conn)
	}
}
