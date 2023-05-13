package routes

import (
	"awesomeProject3/pkg/view"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Router := mux.NewRouter()
	//Router.HandleFunc("/login", view.).Methods("POST")
	Router.HandleFunc("/auth/new-account", view.Registration).Methods("POST")
	Router.HandleFunc("/auth/new-token", view.GenerateNewToken).Methods("POST")
	Router.HandleFunc("/auth/test", view.AuthTest).Methods("POST", "GET")
	Router.HandleFunc("/get-users", view.GetUsers).Methods("POST")
	Router.HandleFunc("/protected", view.HandleProtectedData).Methods("POST")
	Router.HandleFunc("/new-folder", view.CreateFolder).Methods("POST")
	Router.HandleFunc("/get-folder", view.GetUserFolder).Methods("POST")
	Router.HandleFunc("/get-file", view.PrintFile).Methods("POST")
	Router.HandleFunc("/add-file", view.AddFileToFolder).Methods("POST")
	Router.Use(view.Authorize)

	//Router.HandleFunc("/my-files", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")

	return Router
}
