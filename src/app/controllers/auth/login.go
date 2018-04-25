package auth

import (
	"net/http"
	"io"
	h "app/helpers"
)

func Login (w http.ResponseWriter, r *http.Request) {
	h.Log("Request Login", "info")
	io.WriteString(w, "auth")
}
