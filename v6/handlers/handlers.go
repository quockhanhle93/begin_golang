package handlers

import (
	"net/http"
	"fmt"
	"begin/v6/model"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := model.Todos{
		model.Todo{
			Name:      "Write presentation",
			Completed: true,
		},
		model.Todo{
			Name:      "Host meetup",
			Completed: false,
		},
	}

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	fmt.Fprintln(w, "Todo show:", todoId)
}
