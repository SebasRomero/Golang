package main

import "fmt"

func changeValue(str *string) { //Changing the value using the pointer in the RAM
	*str = "Changed!"
}

func changeValue2(str string) string { //if you wanna change the value you should return something
	str = "Changed!2"
	return str
}

func main() {
	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange) //Calling the function to change the original value
	fmt.Println(toChange)
	toChange = changeValue2(toChange) // Returning into a variable the new value
	fmt.Println(toChange)
}
