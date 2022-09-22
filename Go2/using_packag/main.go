package main

import (
	"example.com/packages/countries"
)

func main() {
	countries.Add(("USA"))
	countries.Add(("Colombia"))
	countries.Add(("Perú"))
	countries.Add(("Panamá"))
	countries.Add(("Spain"))
	countries.Add(("Uruguay"))

	err := countries.SetMyCountry("USA")
	if err != nil {
		panic(err)
	}
	countries.List()

}
