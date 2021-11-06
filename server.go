package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

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

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(rw, result)
		}
	})

	handler := cors.Default().Handler(serveMux)
	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            true,
	})

	handler = corHandler.Handler(handler)

	http.ListenAndServe(":5000", handler)
}
