package handlers

import (
	"net/http"
	"fmt"
	"begin_golang/v6/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := model.Todos{
		model.Todo{
			Name:      "Write presentation",
			Completed: true,
			Due: time.Now(),
		},
		model.Todo{
			Name:      "Host meetup",
			Completed: false,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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
