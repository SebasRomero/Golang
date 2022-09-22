package main

import "fmt"

func main() {

	var a []int = []int{1, 7, 4, 5, 6, 8, 7, 3, 8, 9}

	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && i > j {
				fmt.Println(element)
				break
			}

		}
	}

	fmt.Println("-----")

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}

}
