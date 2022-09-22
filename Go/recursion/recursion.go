package main

import "fmt"

/*func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}*/

func factorialRecursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorialRecursion((x - 1))
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(factorialRecursion((4)))
}
