package auth

import (
	"net/http"
	"io"
	h "app/helpers"
)

func Forgot (w http.ResponseWriter, r *http.Request) {
	h.Log("Request Forgot", "info")
	io.WriteString(w, "forgot")
}
