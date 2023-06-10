
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// "github.com/newlinedeveloper/go-api/Configs"

	"github.com/newlinedeveloper/go-api/Routes"


)

func getMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": "Golang project setup test"})

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getMessage).Methods("GET")

	
	// Enable CORS using rs/cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	// Imported Members routes
	Routes.MemberRoutes(router)

	handler := c.Handler(router)

	fmt.Print("Server is running on port 8000 !!!")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
