package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"./entities"
	"github.com/gorilla/mux"
)

var ListPerson []entities.Person

func main() {
	ListPerson = append(ListPerson, entities.NewPerson("Rafael", "Santanna"))
	ListPerson = append(ListPerson, entities.NewPerson("Tatiane", "Alboreli"))

	router := mux.NewRouter()
	router.HandleFunc("/contato", List).Methods("GET")
	router.HandleFunc("/contato/{id}", Get).Methods("GET")
	router.HandleFunc("/contato", Create).Methods("POST")
	router.HandleFunc("/contato/{id}", Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ListPerson)
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range ListPerson {
		Id, error := strconv.ParseInt(params["id"], 10, 64)

		if error == nil {
			fmt.Printf("%d of type %T", Id, Id)
		}
		if item.Id == int(Id) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(500)
	json.NewEncoder(w).Encode("")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range ListPerson {
		Id, error := strconv.ParseInt(params["id"], 10, 64)

		if error == nil {
			fmt.Printf("%d of type %T", Id, Id)
		}

		if item.Id == int(Id) {
			ListPerson = append(ListPerson[:index], ListPerson[index+1:]...)
			break
		}
		json.NewEncoder(w)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var person entities.Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	ListPerson = append(ListPerson, person)
	json.NewEncoder(w).Encode(person)
}
