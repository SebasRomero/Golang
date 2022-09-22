package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string { //Founded in StackOverflow
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // We assigned i=0, j=len(string) and we starting discounting and counting looping until the condition fails
		// fmt.Println(i, j)
		// fmt.Println(runes[i], runes[j])
		runes[i], runes[j] = runes[j], runes[i] //Switching the places in runes
		// fmt.Println(runes[i], runes[j])
	}
	return string(runes)
}

func reverse2(str string) string {
	var res string
	for i := len(str) - 1; i >= 0; i-- {
		res = res + string(str[i])
	}
	return res
}

func reverse3(str string) string {
	var sb strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}

func reverse4(str string) string {
	var res string
	for i := 0; i < len(str); i++ {
		res = string(str[i]) + res
	}
	return res
}

func reverse5(str string) string { //This and the first one are the best solutions, but all are ok
	var res string
	for _, r := range str {
		res = string(r) + res
	}
	return res
}
func main() {
	fmt.Println(reverse("Hello world"))
	fmt.Println(reverse2("Hello world"))
	fmt.Println(reverse3("Hello world"))
	fmt.Println(reverse4("Hello world"))
	fmt.Println(reverse5("Hello world"))
}
