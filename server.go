package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Class struct {
	Check        string `json:"check"`
	Course_title string `json:"course_title"`
	Order        string `json:"order"`
	Title        string `json:"title"`
	Week_order   string `json:"week_order"`
	Week         string `json:"week"`
	Link         string `json:"link"`
	Playtime     string `json:"playtime"`
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("Sparata").Collection("webclass")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var arr bson.M
			if err = cursor.Decode(&arr); err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(rw, arr)
		}

	})

	http.ListenAndServe(":5000", nil)
}
