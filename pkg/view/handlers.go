package view

import (
	"fmt"
	"net/http"
	"text/template"
)

func Handler(_template string, file string, w http.ResponseWriter, r *http.Request) {
	result := _template
	t, err := template.ParseFiles("templates/" + file + ".html")
	if err != nil {
		fmt.Fprintf(w, "error processing")
		return
	}

	tpl := template.Must(t, err)
	tpl.Execute(w, result)
}


