package main

import (
	"fmt"
	"math/big"
)

func NumInList(list []int, num int) bool {
	for _, val := range list { // looping the list to confirm that num there
		if val == num {
			return true
		}
	}
	return false
}

func EvenNumInList(list []int) []int {
	newSlice := []int{}
	for _, val := range list {
		if val%2 == 0 {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func OddNumInList(list []int) []int {
	newSlice := []int{}
	for _, val := range list {
		if val%2 != 0 {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func PrimeNumbers(list []int) []int {
	newSlice := []int{}
	for _, val := range list {
		if big.NewInt(int64(val)).ProbablyPrime(0) { //This returns true if the number inside is prime
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func SumNumberList(list []int) int {
	sum := 0
	for _, val := range list {
		sum += val
	}
	return sum
}

func SumRecursive(list []int) int {
	if len(list) == 0 { //If the list is empty, return 0
		return 0
	}
	return list[0] + SumRecursive(list[1:]) //The first num in the list it's gonna be added to the next first number in the new list starting in the next index later until the list be 0 or empty
}

func FindTwoNumbers(list []int, num int) (int, int) {
	// maxCount := 0
	// indexA, indexB := 0, 0
	for i, val := range list {
		for j, val2 := range list {
			if val+val2 == num && i != j {
				return i, j
			}
		}
	}
	return -1, -1

}

func main() {
	// newSlice := []int{1, 2, 3, 4, 5}
	// fmt.Println(NumInList([]int{1, 2, 3, 4, 5}, 8))
	// fmt.Println(EvenNumInList([]int{1, 2, 3, 6, 8, 4, 5}))
	// fmt.Println(OddNumInList([]int{1, 2, 3, 6, 8, 4, 5}))
	// fmt.Println(PrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	// fmt.Println(SumNumberList([]int{1, 2, 3, 4, -1, -2}))
	// fmt.Println(SumNumberList([]int{}))
	// fmt.Println(SumRecursive([]int{1, 2, 3}))
	// fmt.Println(newSlice[1:])
	fmt.Println(FindTwoNumbers([]int{-2, 9, 1, 8, 2, 3, 4}, 7))
}
