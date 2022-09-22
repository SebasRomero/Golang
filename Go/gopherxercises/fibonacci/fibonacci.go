package main

import "fmt"

func fibonacci(number int) {
	mySlice := []int{1, 1}
	if number == 1 {
	} else {
		for i := 1; i < len(mySlice); i++ {
			mySlice = append(mySlice, mySlice[i]+mySlice[i-1])
			if number == mySlice[i]+mySlice[i-1] {
				break
			}
		}
	}
	fmt.Println("::", mySlice)
}

func main() {
	fibonacci(21)
}
