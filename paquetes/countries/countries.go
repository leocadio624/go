package countries

import (
	"errors"
	"fmt"
)

var myCountry string = "Japan"
var collection []string
var errNotFound error = errors.New("country not found")

// Agrega el pais a la coleccion
func Add(country string) {
	collection = append(collection, country)
}

// Setea variable pais
func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errNotFound
	}
	myCountry = country
	return nil

}

func isInCollection(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}
	return false
}

// Muestra la collecion de paises
func List() {
	for i, c := range collection {
		myCountryMsg := ""
		if c == myCountry {
			myCountryMsg = "[my country]"
		}
		fmt.Printf("%d . %s %s\n", i, c, myCountryMsg)
	}
}
