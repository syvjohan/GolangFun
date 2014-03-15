package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func differenses(str, str2 string) string {

	dmp := diffmatchpatch.Diff(str, str2)

	return dmp
}

func main() {
	var dstr string = "hello"
	var str2 string = "hello"

	fmt.Println(differenses(str, str2))
}
