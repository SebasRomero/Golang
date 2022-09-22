package main

import "fmt"

func test(x int) {
	fmt.Println("Hello!", x)
}

func main() {
	/*x := test
	x(5)*/
	test2 := func(x int) int {
		return x * -1
	}(9)
	fmt.Println(test2)

}
