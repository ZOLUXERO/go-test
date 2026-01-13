Ejemplo de como leer y escribir json en go
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
    // Ejemplo json
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	// Convierta un json a una estructura para poder leer el json que llega
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Println(`unmarshalled: `, unmarshalled)

	// Convierta una estructura a json
	var mySlice []Person
	var m1 Person
	m1.FirstName = "wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		log.Println("Error marshaling", err)
	}
	fmt.Println(string(newJson))
}
```
