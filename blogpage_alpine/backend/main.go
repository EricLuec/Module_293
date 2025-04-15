package main

import (
	"./controllers"
	"./routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	controllers.Client = client

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/").Handler(fs)

	fmt.Println("Server läuft auf http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
