package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	var text1 string = "hejj"
	var text2 string = "hej"
	var dmp = diffmatchpatch.New()

	var diffs = dmp.DiffMain(text1, text2, false)

	fmt.Println("%s", diffs)
}
