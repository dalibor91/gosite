package index

import (
	"app/controllers"
	h "app/helpers"
	p "app/properties"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Time (w http.ResponseWriter, r *http.Request) {
	controllers.SetJsonHeadersOk(w, fmt.Sprintf(`{"time": "%s"}`, h.GetTimestamp()))
}

func Date (w http.ResponseWriter, r *http.Request) {
	controllers.SetJsonHeadersOk(w, fmt.Sprintf(`{"date": "%s"}`, h.GetDate()))
}

func Login (w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile(fmt.Sprintf("%s/templates/main/login.html", p.Get("server.host.static_dir")))
	w.Write(body)
}

func Register (w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile(fmt.Sprintf("%s/templates/main/register.html", p.Get("server.host.static_dir")))
	w.Write(body)
}

func Forgot (w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile(fmt.Sprintf("%s/templates/main/forgot.html", p.Get("server.host.static_dir")))
	w.Write(body)
}
