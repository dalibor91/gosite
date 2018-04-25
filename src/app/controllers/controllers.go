package controllers

import (
	"net/http"
	"io"
)

func SetJsonHeadersOk(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, response)
}

func SetHTMLHeadersOk(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, response)
}

func GetAuthorization(r *http.Request) {

}




