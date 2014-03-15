package main

import (
	"fmt"
)

func strcmp(a, b string) int {
	var min = len(b)
	if len(a) < len(b) {
		min = len(a)
	}

	var diff int
	for i := 0; i < min && diff == 0; i++ {
		diff = int(a[i]) - int(b[i])
	}
	if diff == 0 {
		diff = len(a) - len(b)
	}
	fmt.Println("diff is: %d", diff)
	return diff
}

func main() {
	var a string = "hej"
	var b string = "hej"

	strcmp(a, b)

}
