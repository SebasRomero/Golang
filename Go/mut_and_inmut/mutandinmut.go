package main

import "fmt"

func changeFirst(slice []int) { // Here changes the value in the 0 position
	slice[0] = 1000
}

func changeLast(slice []int) { // Here changes the value in the last position
	slice[len(slice)-1] = 1000
}
func main() {
	var x []int = []int{3, 4, 5, 7, 9} // Create a slice
	fmt.Println(x)
	changeFirst(x) // Calling this method sending the slice as an argument
	changeLast(x)  // The same with another function
	fmt.Println(x)
}
