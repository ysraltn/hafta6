package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Input struct {
	Start string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(responseWriter, nil)
			return
		}

		details := Input{
			Start: request.FormValue("start"),
		}
		if details.Start == "start" {
			fmt.Println(details.Start)
			fmt.Println(details)
		}

		tmpl.Execute(responseWriter, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8085", nil)
}
