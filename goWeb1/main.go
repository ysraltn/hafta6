package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students/{id}", func(responseWriter http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		fmt.Fprintf(responseWriter, "student with id: %s \n", id)
	})

	http.ListenAndServe(":85", r)
}
