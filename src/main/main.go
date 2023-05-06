package main

import (
	"awesomeProject3/pkg/routes"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server started")
	n := &http.Server{
		Handler:      routes.Route(),
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := n.ListenAndServe()
	if err != nil {
		return
	}
}


