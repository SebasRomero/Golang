package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 7,
	}

	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(len(mp))

}
