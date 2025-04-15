package main

import (
	"blogapline/controllers"
	"blogapline/routes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

// CORS Middleware
func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	// MongoDB verbinden
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Fehler beim Erstellen des MongoDB-Clients:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Fehler beim Verbinden mit MongoDB:", err)
	}

	controllers.Client = client

	// Router konfigurieren
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Statische Dateien (HTML, CSS, JS) aus ./public ausliefern
	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/").Handler(fs)

	// CORS aktivieren und Server starten
	fmt.Println("✅ Server läuft auf http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", enableCORS(router)))
}
