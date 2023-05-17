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
	Router.HandleFunc("/auth/validate-token", view.ValidateTokenC).Methods("POST")
	Router.HandleFunc("/auth/tests", view.AuthTest).Methods("POST", "GET")
	Router.HandleFunc("/drive/upload", view.UploadFileHandler).Methods("POST")
	Router.HandleFunc("/drive/delete", view.DeleteUserFile).Methods("POST")
	Router.HandleFunc("/drive/get-folder", view.GetUserFolder).Methods("POST")
	Router.HandleFunc("/drive/new-folder", view.CreateFolder).Methods("POST")
	Router.HandleFunc("/drive/rename", view.RenameUserFile).Methods("POST")
	Router.HandleFunc("/drive/download", view.DownloadFileHandler).Methods("POST", "GET")
	Router.HandleFunc("/drive/tests", view.TestDownload).Methods("POST", "GET")
	Router.HandleFunc("/get-user", view.GetUserData).Methods("POST")
	Router.HandleFunc("/get-users", view.GetUsers).Methods("POST")
	Router.HandleFunc("/protected", view.HandleProtectedData).Methods("POST")
	Router.HandleFunc("/get-file", view.PrintFile).Methods("POST")
	//Router.HandleFunc("/add-file", view.AddFileToFolder).Methods("POST")
	Router.Use(view.Authorize)

	//Router.HandleFunc("/my-files", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")
	//Router.HandleFunc("/files/{folderAccessLink}", view.TestView).Methods("POST")

	return Router
}
