package routes

import (
	"awesomeProject3/pkg/view"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Router := mux.NewRouter()
	//Router.HandleFunc("/login", view.).Methods("POST")
	Router.HandleFunc("/new-account", view.Registration).Methods("POST")
	Router.HandleFunc("/get-users", view.GetUsers).Methods("POST")
	Router.HandleFunc("/protected", view.HandleProtectedData).Methods("POST")
	Router.Use(view.Authorize)
	Router.HandleFunc("/new-token", view.GenerateNewToken).Methods("POST")
	//Router.HandleFunc("/my-files", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")

	return Router
}
