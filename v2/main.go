package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
	"github.com/gorilla/mux"
)

func main() {
	router:= mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))

}
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello, %q", html.EscapeString(request.URL.Path))
}