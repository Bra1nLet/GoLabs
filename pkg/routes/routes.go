package routes

import (
	"awesomeProject3/pkg/view"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", view.TestView).Methods("GET")
	Router.HandleFunc("/login", view.TestView).Methods("GET")
	Router.HandleFunc("/register", view.TestView).Methods("GET")
	Router.HandleFunc("/my-files", view.TestView).Methods("GET")
	Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("GET")
	return Router
}
