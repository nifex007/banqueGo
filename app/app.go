package app

import (
	"fmt"
	"log"
	"net/http"
)


func Start() {

	port := ":8081"
	mux := http.NewServeMux()
	// Routes
	mux.HandleFunc("/saluer", saluer)
	mux.HandleFunc("/clients", listerLesClients)
	//  starts server in port 8081
	fmt.Println("localhost @ "+port)
	log.Fatal(http.ListenAndServe("localhost"+port, mux))

}