package view

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data

	err := r.ParseMultipartForm(10 << 20) // Max file size: 10MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := generateFileName(handler.Filename)
	ff, _ := handler.Open()
	AddFileToFolder(ff, filename, w, r)
	fmt.Fprintf(w, "File uploaded successfully")
}

func generateFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	filename := filepath.Base(originalName[:len(originalName)-len(ext)])
	return filename + "_" + generateRandomString(6) + ext
}

func generateRandomString(length int) string {
	// Implement your own random string generation logic here
	// This is just a placeholder implementation
	return "random"
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the file name from the request parameters
	vars := mux.Vars(r)
	filename := vars["filename"]

	// Specify the file path where the file is located
	// In this example, the files are assumed to be in the /client folder
	filepath := filepath.Join(".", "client", filename)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate headers for the file download
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	// Stream the file to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error streaming file", http.StatusInternalServerError)
		return
	}
}
