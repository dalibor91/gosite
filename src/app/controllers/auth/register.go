package auth

import (
	"net/http"
	"io"
	h "app/helpers"
)

func Register (w http.ResponseWriter, r *http.Request) {
	h.Log("Request Register", "info")
	io.WriteString(w, "register");
}
