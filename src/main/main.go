package main

import (
	"awesomeProject3/models"
	"awesomeProject3/routes"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server started")
	models.Migrate()
	n := &http.Server{
		Handler:      routes.Route(),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := n.ListenAndServe()
	if err != nil {
		return
	}
}
