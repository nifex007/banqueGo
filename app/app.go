package app

import (
	"fmt"
	"log"
	"net/http"
)


func start() {

	port := ":8081"
	// Routes
	http.HandleFunc("/saluer", saluer)
	http.HandleFunc("/clients", listerLesClients)
	//  starts server in port 8081
	fmt.Println("localhost @ "+port)
	log.Fatal(http.ListenAndServe("localhost"+port, nil))

}