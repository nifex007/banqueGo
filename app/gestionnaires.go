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



func saluer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Bonjour le monde!")
}

func listerLesClients(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Name: "John", City: "Paris", Zipcode: "75008"},
		{Name: "Mary", City: "London", Zipcode: "W1A 0AX"},
	}
	jsonHeader := "application/json"
	xmlHeader := "application/xml"
	
	if request.Header.Get("Content-Type") ==  xmlHeader {
		writer.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	}else{
		writer.Header().Set("Content-Type", jsonHeader)
		json.NewEncoder(writer).Encode(customers)
	}
}

