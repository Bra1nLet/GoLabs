package view

import (
	"net/http"
)

func TestView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	Handler("test", "files", w, r)
}
