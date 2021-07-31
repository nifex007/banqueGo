package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string	`json:"full_name"`
	City    string	`json:"city"`
	Zipcode string	`json:"zipcode"`
}

func main() {
	// Routes
	http.HandleFunc("/saluer", saluer)
	http.HandleFunc("/clients", listerLesClients)
	//  starts server in port 8081
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func saluer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Bonjour le monde!")
}

func listerLesClients(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Name: "John", City: "Paris", Zipcode: "75008"},
		{Name: "Mary", City: "London", Zipcode: "W1A 0AX"},
	}

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(customers)
}
