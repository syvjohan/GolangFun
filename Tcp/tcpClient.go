package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Struct struct {
	M, N int64
}

func main() {
	fmt.Println("Start the client...");
	conn, err := net.Dial("tcp", "localhost:8080") //connects to the address on the named network.
	if err != nil {
		log.Fatal("Connection error...", err)

	}
	encoder := gob.NewEncoder(conn) // NewEncoder returns a new encoder that will transmit on the io.Writer.
	pointer := &Struct{1, 2}
	encoder.Encode(pointer)
	conn.Close()
	fmt.Println("done");
}
