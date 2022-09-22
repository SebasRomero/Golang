package main

import "fmt"

func test(x, y float64) (ret1, ret2, ret3, ret4 float64) {
	defer fmt.Println("Has been returned with success") //Shows success for testing
	defer fmt.Println("This prints first")
	ret1 = x + y
	ret2 = x - y
	ret3 = x * y
	ret4 = x / y
	return
}

func main() {
	plus, subs, multiply, divide := test(9, 1)
	fmt.Println("Plus: ", plus,
		"Subs: ", subs, "Multiply: ",
		multiply, "Divide: ", divide)
}
