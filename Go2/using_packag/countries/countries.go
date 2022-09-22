package countries

import (
	"errors"
	"fmt"
)

var MyCountry string = ""
var collection []string
var errNotFound error = errors.New("Country not found")

// Add a country in the collection
func Add(country string) {
	collection = append(collection, country)
}

// Change my country
func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errNotFound
	} else {
		MyCountry = country
		return nil
	}
}

// Loop the collection
func List() {
	for i, val := range collection {
		myCountryMsg := ""
		if val == MyCountry {
			myCountryMsg = "[My country]"
		}
		fmt.Printf("%d.%s %s\n", i+1, val, myCountryMsg)
	}
}

// isInCollection finds if that country exists
func isInCollection(country string) bool {
	for _, val := range collection {
		if val == country {
			return true
		}
	}
	return false
}
