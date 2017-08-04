package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"begin/v5/model"
	"encoding/json"
	"log"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(writer, "Todo show:", todoId)
}
func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	todos := model.Todos{
		model.Todo{Name:"Write presentation"},
		model.Todo{Name:"Host meetup"},
	}

	if error := json.NewEncoder(writer).Encode(todos); error != nil{
		panic(error)
	}
}
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome!")
}
