package main

import "fmt"

func multiply(a, b float64) (c float64) {

	c = a * b

	fmt.Println("Calculated value is: ", c)
	return
}

func division(a, b float64) (c float64) {

	c = a / b

	fmt.Println("Calculated value is: ", c)
	return
}

func addition(a, b float64) (c float64) {

	c = a + b

	fmt.Println("Calculated value is: ", c)
	return
}

func subtract(a, b float64) (c float64) {

	c = a + b

	fmt.Println("Calculated value is: ", c)
	return
}

func cases(choice int, a, b float64) {

	switch choice {

	case 0:
		choice = 0
		division(a, b)

	case 1:
		choice = 1
		multiply(a, b)

	case 2:
		choice = 2
		subtract(a, b)

	case 3:
		choice = 3
		addition(a, b)

	default:
		fmt.Println("Must choice one of the four")
	}

	return
}

func menu() {
	fmt.Println("Please select one of the following")
	fmt.Println("Press 0 for division")
	fmt.Println("Press 1 for multiply")
	fmt.Println("Press 2 for subtract")
	fmt.Println("Press 3 for addition")
}

func main() {

	menu()

	var choice int
	var a, b float64

	n, err := fmt.Scanf("%d\n", &choice)

	if err != nil && n != 1 {
		fmt.Println(n, err)
	}

	fmt.Println("please enter a first value")

	p, error := fmt.Scanf("%f\n", &a)
	if error != nil && p != 1 {
		fmt.Println(p, error)
	}
	fmt.Println("first value: ", a)

	fmt.Println("please enter a second value")

	k, erro := fmt.Scanf("%f\n", &b)
	if erro != nil && k != 1 {
		fmt.Println(k, erro)
	}
	fmt.Println("second value: ", b)

	cases(choice, a, b)
}
