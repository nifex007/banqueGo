package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func main() {
	port := ":8081"
	// Routes
	http.HandleFunc("/saluer", saluer)
	http.HandleFunc("/clients", listerLesClients)
	http.HandleFunc("/xml/clients", listerLesClientsXML)
	//  starts server in port 8081
	fmt.Println("localhost @ "+port)
	log.Fatal(http.ListenAndServe("localhost"+port, nil))

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

func listerLesClientsXML(writer http.ResponseWriter, request *http.Request){
	customers := []Customer{
		{Name: "John", City: "Paris", Zipcode: "75008"},
		{Name: "Mary", City: "London", Zipcode: "W1A 0AX"},
	}

	writer.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(writer).Encode(customers)
}
