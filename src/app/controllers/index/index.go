package index

import (
	"app/controllers"
	h "app/helpers"
	"net/http"
	"fmt"
)

func Time (w http.ResponseWriter, r *http.Request) {
	h.Log("Request Time", "info")
	controllers.SetJsonHeadersOk(w, fmt.Sprintf(`{"time": "%s"}`, h.GetTimestamp()))
}

func Date (w http.ResponseWriter, r *http.Request) {
	h.Log("Request Date", "info")
	controllers.SetJsonHeadersOk(w, fmt.Sprintf(`{"date": "%s"}`, h.GetDate()))
}

func Index(w http.ResponseWriter, r *http.Request) {
	h.Log("Request at index")
}