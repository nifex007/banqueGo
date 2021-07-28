package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/saluer", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Bonjour le monde!")
	})

	http.ListenAndServe("localhost:8081", nil)
}
