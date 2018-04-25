package server

import (
	mux "github.com/gorilla/mux"
	ctrlauth "app/controllers/auth"
	crelindex "app/controllers/index"
	"strings"
	"net/http"
	"fmt"
	h "app/helpers"
)

func main(r * mux.Router, static_location string, static_url string) (* mux.Router) {
	r.HandleFunc("/", ctrlauth.Login).Methods("GET", "POST")
	r.HandleFunc("/server.time", crelindex.Time).Methods("GET", "POST")
	r.HandleFunc("/server.date", crelindex.Date).Methods("GET", "POST")

	if strings.Trim(static_location, " ") != "" && static_url != "" {
		h.Log(fmt.Sprintf("Static location: %s", static_location), "debug")
		h.Log(fmt.Sprintf("Static url: %s", static_url), "debug")
		r.PathPrefix(static_url).Handler(http.StripPrefix(static_url, http.FileServer(http.Dir(static_location))))
	}

	return r
}

func post(r * mux.Router) (* mux.Router) {
	h.Log("setting POST routes", "debug")
	return r
}

func get(r * mux.Router) (* mux.Router) {
	h.Log("setting GET routes", "debug")

	r.HandleFunc("/auth", ctrlauth.Login).Methods("GET")
	r.HandleFunc("/register", ctrlauth.Register).Methods("GET")
	r.HandleFunc("/forgot", ctrlauth.Forgot).Methods("GET")
	return r
}

func Routes(static_location string, static_url string) * mux.Router {
	return post(get(main(mux.NewRouter(), static_location, static_url)))
}


