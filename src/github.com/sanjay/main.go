package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/handlers"

	"github.com/models"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBNAME Database name
const DBNAME = "phonebook"

// COLLECTION Collection name
const COLLECTION = "people"

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://root:password@mongo:27017"

func init() {
	// Populates database with dummy data

	var people []models.Person

	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(DBNAME)

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("person_data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValues, &people)

	// Insert people into DB
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}
	_, err = db.Collection(COLLECTION).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", handlers.GetAllPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", handlers.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}", handlers.UpdatePersonEndpoint).Methods("PUT")
	log.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
