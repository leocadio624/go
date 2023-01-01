package main

import "example.com/pequetes/countries"

func main() {
	//id := uuid.New()
	//fmt.Println(id)

	countries.Add("Mexico")
	countries.Add("EUA")
	countries.Add("Canada")
	countries.Add("Spain")

	err := countries.SetMyCountry("Mexico")
	if err != nil {
		panic(err)
	}

	countries.List()

}
