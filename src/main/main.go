package main

import (
	"awesomeProject3/pkg/routes"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server started")
	n := &http.Server{
		Handler: routes.Route(),
		Addr:    "127.0.0.1:8080"}
	err := n.ListenAndServe()
	if err != nil {
		return
	}
}
