package main

import "fmt"

func randomValue(input int) (calc int) {
	const bitsPerWord = 32 << uint(^uint(0)>>63)
	const BitsPerWord = bitsPerWord
	maxInt := 1<<(bitsPerWord-1) - 1

	var primSmal int = 28657
	var primBig int = 214229

	var bitsh = -(input << 16)

	calc = bitsh * (primBig + primSmal) * maxInt

	return

}

func GenerateRandomValue(input int) {
	slice := make([]int, 0)
	var f int

	for i := 0; i != 1000; i++ {
		f += randomValue(input)
		slice = append(slice, f)
	}

	for i := range slice {
		fmt.Println(i, slice[i])
	}

}

func main() {
	menu()
	handleInput()

}

func handleInput() (input int) {

	p, error := fmt.Scanf("%d\n", &input)
	if error != nil && p != 1 {
		fmt.Println(p, error)
	}

	if (1 <= input) && (input <= 99999) {
		GenerateRandomValue(input)
	} else {
		fmt.Println("Input value need to be between 1 and 99999")
	}

	return
}

func menu() {
	fmt.Println("Welcome to the random generator program")
	fmt.Println("-----------------------------------------------------\n")
	fmt.Println("Please print a value to generato random value from: ")
}
