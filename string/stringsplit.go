package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func stringsplit(str string) (result string) {

	words := strings.Split(str, ".")
	ip, port := words[0], words[1]
	if ip != "" && port != "" {
		fmt.Println(ip, port)
	}

	return result
}

func readfile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text := string(file[:])

	return text
}

func writefile(path string, result []byte) {
	err := ioutil.WriteFile(path, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	stringsplit(readfile("content.txt"))
	//textwrite := writefile("container")

}
