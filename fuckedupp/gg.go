package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"net/http"
	"Lib/checksum"
)

// Echo the data received on the WebSocket.
func echoServer(ws *websocket.Conn) {
clientShadow := shadow{shadowstring:"shadow"}
clientBackup := backup{backupstring:"backup"}
	for {
		buffer := make([]byte, 512)
		readData, _ := ws.Read(buffer)
		str := string(buffer[:readData])

		if str != "" {
			differances(str, &clientBackup, &clientShadow)

			readfile("fileServer.txt")
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

func differances(text string, clientBackup *backup, clientShadow *shadow) {
	clientShadowLocalVersion := shadow{shadowLocalVersion: 0}
	//Check diffs
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(clientShadow.shadowstring, text, false)
	if len(diffs) == 1 {
		// no diffs
		return
	}
	dmp.DiffCleanupEfficiency(diffs)
	// Make patch
	patch := dmp.PatchMake(clientShadow.shadowstring, diffs)

	fmt.Println(dmp.PatchToText(patch))
	result := dmp.PatchToText(patch)

	clientShadow.shadowstring = text

	clientShadowLocalVersion.shadowLocalVersion++

	/*// Check diffs
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
	*/ 
}

func sendEdits() {

}

func sendAck() {

}

func handleEditMessage() {

}

func sendRequest() {
	
}

//<-------------------------------------------------->




//<----------------------------------------------------------------------->

type shadow struct {
	shadowLocalVersion  int
	shadowRemoteVersion int
	shadowstring        string
}

type backup struct {
	backupstring           string
	backupLocalVersion int
}

type Client struct {
	client int
}

func main() {
	fmt.Println("Initialize server...")
	connectionHandler()
}
