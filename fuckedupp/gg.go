package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"net/http"
)

// Echo the data received on the WebSocket.
func echoServer(ws *websocket.Conn) {

	for {
		buffer := make([]byte, 512)
		readData, _ := ws.Read(buffer)
		str := string(buffer[:readData])

		if str != "" {
			differances(readfile("serverShadow.txt"), str)

			result := readfile("resultClient.txt")
			if result != "" {
				bytes := []byte(result)
				ws.Write(bytes)
			}
		}
	}
}

func connectionHandler() {
	http.Handle("/", websocket.Handler(echoServer))
	fmt.Println("Start listening after clients...")

	err := http.ListenAndServe(":12345", nil)
	// if err not equal null then panic (stop execution)
	if err != nil {
		panic("Error ListenAndServe: " + err.Error())
	}
}

//<---------------------------------------------------------------->

func readfile(path string) string {
	inbyte, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text := string(inbyte[:])

	return text
}

func differances(text1, text2 string) {
	// Check diffs
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)
	// Make patch
	patch := dmp.PatchMake(text1, diffs)

	// Print patch to see the results before they're applied
	fmt.Println(dmp.PatchToText(patch))

	result := dmp.PatchToText(patch)

	// detta sker efter att vi har mottagit från klienten!
	// Apply patch
	//text3, _ := dmp.PatchApply(patch, text1)

	// Write the patched text to a new file
	err := ioutil.WriteFile("resultClient.txt", []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}

//<----------------------------------------------------------------------->

func main() {
	fmt.Println("Initialize server...")
	connectionHandler()
}
