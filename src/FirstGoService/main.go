package main

import (
	"encoding/json"
	"fmt"
	"log"

	"./entities"
)

var peaple []entities.Person

func main() {
	peaple = append(peaple, entities.NewPerson(1, "Rafael", "Santanna"))
	peaple = append(peaple, entities.NewPerson(2, "Tatiane", "Alboreli"))

	resp, err := json.Marshal(peaple)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Printf(string(resp))
}
