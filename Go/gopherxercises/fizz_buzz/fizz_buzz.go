package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(number int) string { //My version
	res := []int{}
	strRes := ""
	for i := 1; i <= number; i++ {
		res = append(res, i)
	}
	for i := 0; i < len(res); i++ {

		if len(res) == 1 {
			strRes = strconv.Itoa(1)
			break
		} else {

			if res[i]%3 == 0 && res[i]%5 == 0 {
				if i+1 == len(res) {
					strRes = strRes + "Fizz Buzz"
				} else {
					strRes = strRes + "Fizz Buzz, "
				}
				continue
			}
			if res[i]%3 == 0 {
				if i+1 == len(res) {
					strRes = strRes + "Fizz"
				} else {
					strRes = strRes + "Fizz, "
				}
				continue
			}
			if res[i]%5 == 0 {
				if i+1 == len(res) {
					strRes = strRes + "Buzz"
				} else {
					strRes = strRes + "Buzz, "
				}
				continue
			}
			if i+1 == len(res) {
				strRes = strRes + strconv.Itoa(res[i])
			} else {
				strRes = strRes + strconv.Itoa(res[i]) + ", "
			}
		}
	}
	return strRes
}

func fizzBuzz2(n int) { //The good way
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
}

func main() {
	// fmt.Print(fizzBuzz(30))
	fizzBuzz2(15)
}
