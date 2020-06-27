package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type aPerson struct {
	// Create a nested json response basically, but in Go, it is called a struct
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *address `json:"address,omitempty"`
}

type address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// This is basically a list of persons, called people
var people []aPerson

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Loops through the people list returning details of the id if found
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// After looping, encodes the response
	json.NewEncoder(w).Encode(&aPerson{})
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	// Retrieves the list of hardcoded people
	json.NewEncoder(w).Encode(people)
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Post request that will create a new person and append it to people list
	params := mux.Vars(req)
	var person aPerson
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Delete uses index to slice out result from the list
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, aPerson{ID: "1", Firstname: "Uche", Lastname: "Alozie",
		Address: &address{City: "Eko", State: "Lagos"}})
	people = append(people, aPerson{ID: "2", Firstname: "Ifeoma", Lastname: "Alozie"})
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5050", router))

}
