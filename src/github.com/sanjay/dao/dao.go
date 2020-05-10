package dao

import (
	"context"
	"log"

	"github.com/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://root:password@mongo:27017"

// DBNAME Database name
const DBNAME = "phonebook"

// COLLNAME Collection name
const COLLNAME = "people"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// InsertManyValues inserts many items from byte slice
func InsertManyValues(people []models.Person) {
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}
	_, err := db.Collection(COLLNAME).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertOneValue inserts one item from Person model
func InsertOneValue(person models.Person) {
	log.Println(person)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllPeople returns all people from DB
func GetAllPeople() []models.Person {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Person
	var elem models.Person
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

// DeletePerson deletes an existing person
func DeletePerson(person models.Person) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), person, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdatePerson updates an existing person
func UpdatePerson(person models.Person, personID string) {
	doc := db.Collection(COLLNAME).FindOneAndReplace(
		context.Background(),
		bson.D{primitive.E{Key: "id", Value: personID}},
		person,
	)

	log.Println(doc)
}
