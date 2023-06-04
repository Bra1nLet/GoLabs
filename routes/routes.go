package routes

import (
	"awesomeProject3/handlers"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", handlers.Hello).Methods("GET")
	Router.HandleFunc("/auth/new-account", handlers.Registration).Methods("POST")
	Router.HandleFunc("/auth/new-token", handlers.GenerateNewToken).Methods("POST")
	Router.HandleFunc("/auth/validate-token", handlers.ValidateTokenC).Methods("POST")
	Router.HandleFunc("/drive/upload", handlers.UploadFileHandler).Methods("POST")
	Router.HandleFunc("/drive/delete", handlers.DeleteUserFile).Methods("POST")
	Router.HandleFunc("/drive/get-folder", handlers.GetUserFolder).Methods("POST")
	Router.HandleFunc("/drive/new-folder", handlers.CreateFolder).Methods("POST")
	Router.HandleFunc("/drive/rename", handlers.RenameUserFile).Methods("POST")
	Router.HandleFunc("/drive/download", handlers.DownloadFileHandler).Methods("POST", "GET")
	Router.HandleFunc("/drive/print-file", handlers.PrintFile).Methods("POST")
	Router.HandleFunc("/get-user", handlers.GetUserData).Methods("POST")
	Router.HandleFunc("/get-users", handlers.GetUsers).Methods("POST")
	Router.HandleFunc("/protected", handlers.HandleProtectedData).Methods("POST")
	Router.Use(handlers.Authorize)
	return Router
}
