package handlers

import (
	"fmt"
	"net/http"
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
	return filename + "_" + ext
}
