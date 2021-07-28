package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/saluer", saluer)
	//  starts server 1n port 8081
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func saluer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Bonjour le monde!")
}