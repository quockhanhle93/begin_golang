package main

import (
	"log"
	"net/http"
	"begin/v6/routes"
)

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
