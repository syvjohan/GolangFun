package main

import (
	"fmt"
)

func strcmp(a, b string) (diff int) {
	var min = len(b)
	if len(a) < len(b) {
		min = len(a)
	}

	for i := 0; i < min && diff == 0; i++ {
		diff = int(a[i]) - int(b[i])
	}
	if diff == 0 {
		diff = len(a) - len(b)
	}

	return diff
}

func main() {
	var a string = "hejj"
	var b string = "hejjj"

	fmt.Println("The diff is: ", strcmp(a, b))

}
