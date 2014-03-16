package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	var text1 string = "hej"
	var text2 string = "hejj"
	var dmp = diffmatchpatch.New()

	var diffs = dmp.DiffMain(text1, text2, false)

	fmt.Println("%s", diffs)
}
