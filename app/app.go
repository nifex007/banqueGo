package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)


func Start() {

	port := ":8081"
	router := mux.NewRouter()
	// Routes
	router.HandleFunc("/saluer", saluer)
	router.HandleFunc("/clients", listerLesClients)
	//  starts server in port 8081
	fmt.Println("localhost @ "+port)
	log.Fatal(http.ListenAndServe("localhost"+port, router))

}