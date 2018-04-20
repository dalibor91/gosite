package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	p "app/properties"
	s "app/server"
)

func setProps() {
	usr, err := user.Current()
	if err != nil {
		log.Panic("Unable to get current user")
		os.Exit(1)
	} else {
		p.Set("app.dir", fmt.Sprintf("%s/.dweb", usr.HomeDir))
	}

	p.Set("app.config", fmt.Sprintf("%s/config.ini", p.Get("app.dir")))

	if _, err := os.Stat(p.Get("app.config")); !os.IsNotExist(err) {
		log.Panic(fmt.Sprintf("file: %s does not exists", p.Get("app.config")))
		os.Exit(1)
	}

	p.Set("server.host.bind", "0.0.0.0:10050")
	p.Set("server.host.name", "www.dalibor.test")
	p.Set("server.host.alias", "dalibor.test")
	p.Set("server.host.static_dir", "/go/api/static")
	p.Set("server.host.static_url", "/assets/")
	p.Set("server.host.template_dir", "/go/api/templates")
	p.Dump()
}

func main() {
	setProps()
	router := s.Routes(p.Get("server.host.static_dir"), p.Get("server.host.static_url"))
	log.Fatal(http.ListenAndServe(p.Get("server.host.bind"), router))
}
