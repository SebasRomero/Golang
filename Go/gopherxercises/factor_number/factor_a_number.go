package main

import "fmt"

func factor(list []int, number int) []int { //My way to do
	mySlice := []int{}
	for number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		if number%2 == 0 {
			number = number / 2
			mySlice = append(mySlice, 2)
		} else if number%3 == 0 {
			number = number / 3
			mySlice = append(mySlice, 3)
		} else if number%5 == 0 {
			number = number / 5
			mySlice = append(mySlice, 5)
		} else if number%7 == 0 {
			number = number / 7
			mySlice = append(mySlice, 7)
		}

	}
	return mySlice
}

func factor2(primes []int, number int) []int { //The optimal way
	var res []int
	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number = number / prime
		}
	}
	if number > 1 {
		res = append(res, number)
	}
	return res
}

func main() {
	fmt.Println(factor([]int{2, 3, 5}, 720))
	fmt.Println(factor2([]int{2, 3, 5}, 720))
}
