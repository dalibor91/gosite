package server

import (
	mux "github.com/gorilla/mux"
	ctrl "app/controllers/index"
	"strings"
	"net/http"
	"log"
	"fmt"
)

func main(r * mux.Router, static_location string, static_url string) (* mux.Router) {
	r.HandleFunc("/", ctrl.Login).Methods("GET", "POST")

	r.HandleFunc("/server.time", ctrl.Time).Methods("GET", "POST")
	r.HandleFunc("/server.date", ctrl.Date).Methods("GET", "POST")

	if strings.Trim(static_location, " ") != "" && static_url != "" {
		log.Println(fmt.Sprintf("Static location: %s", static_location))
		log.Println(fmt.Sprintf("Static url: %s", static_url))
		r.PathPrefix(static_url).Handler(http.StripPrefix(static_url, http.FileServer(http.Dir(static_location))))
	}

	return r
}

func post(r * mux.Router) (* mux.Router) {
	return r
}

func get(r * mux.Router) (* mux.Router) {
	r.HandleFunc("/login", ctrl.Login).Methods("GET")
	r.HandleFunc("/register", ctrl.Register).Methods("GET")
	r.HandleFunc("/forgot", ctrl.Register).Methods("GET")
	return r
}

func Routes(static_location string, static_url string) * mux.Router {
	return post(get(main(mux.NewRouter(), static_location, static_url)))
}


