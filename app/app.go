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
	router.HandleFunc("/saluer", saluer).Methods(http.MethodGet)
	router.HandleFunc("/clients", listerLesClients).Methods(http.MethodGet)
	
	router.HandleFunc("/client/{client_id:[0-9]+}", obtenirUnClient).Methods(http.MethodGet)
	router.HandleFunc("/clients", créerUnClient).Methods(http.MethodPost)
	//  starts server in port 8081
	fmt.Println("localhost @ "+port)
	log.Fatal(http.ListenAndServe("localhost"+port, router))

}

func obtenirUnClient(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	client_id := vars["client_id"]
	fmt.Println("client_id: " + client_id)
	fmt.Fprintf(writer, "client_id: %s", client_id)
}

func créerUnClient(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello")
}