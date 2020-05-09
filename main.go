package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person struct
type Person struct {
	ID          string `json:"id,omitempty"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	ContactInfo `json:"contactinfo,omitempty"`
}

// ContactInfo Contact info of the person
type ContactInfo struct {
	City    string `json:"city,omitempty"`
	ZipCode string `json:"zipcode,omitempty"`
	Phno    string `json:"phone,omitempty"`
}

var people []Person

// AddPerson new
func AddPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

// GetPeoples Get all peoples
func GetPeoples(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)

}

// GetPerson return requested person
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, p := range people {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
		json.NewEncoder(w).Encode("Person not found")
	}
}

// UpdatePersonDetails Update  update
func UpdatePersonDetails(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			people[i] = person
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}

// DeletePerson from the record
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, p := range people {
		if p.ID == params["id"] {
			fmt.Println("id:", params["id"])
			copy(people[i:], people[i+1:])
			people = people[:len(people)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", AddPerson).Methods("POST")
	router.HandleFunc("/peoples", GetPeoples).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", UpdatePersonDetails).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	// Trigger server
	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
