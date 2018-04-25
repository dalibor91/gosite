package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	p "app/properties"
	s "app/server"
	h "app/helpers"
)

func setProps() {
	h.Log("main.setProps()", "debug")

	usr, err := user.Current()
	if err != nil {
		h.Log("Unable to get current user", "error")
		os.Exit(1)
	} else {
		app_dir := fmt.Sprintf("%s/.dweb", usr.HomeDir)
		h.Log(fmt.Sprintf("main.setProps() -> HomeDir=%s", usr.HomeDir), "debug")
		p.Set("app.dir", app_dir)
		if _, err := os.Stat(app_dir); err != nil {
			fmt.Println("not found ", app_dir)
			os.Mkdir(app_dir, 700)
		}
	}

	user_config := fmt.Sprintf("%s/config.ini", p.Get("app.dir"))
	p.Set("app.config", user_config)

	if _, err := os.Stat(user_config); err != nil {
		h.Log(fmt.Sprintf("Config file '%s' does not exists or can not be read", user_config))
		os.Exit(1)
	}

	ini, errini := h.NewINIReader(user_config)

	option_hostbind := "0.0.0.0:10050"
	option_hostname := "dalibor.me"
	option_hostalias := "www.dalibor.me"
	option_staticdir := "/go/api/static"
	option_staticurl := "/assets/"
	option_templatedir := "/go/api/templates"

	if errini == nil {
		h.Log("ini found read ini", "debug")
		if sec, errsec := ini.GetIni().GetSection("host"); errsec == nil {
			h.Log("read section 'host'", "debug")
			if _host, err := sec.GetKey("bind"); err == nil { option_hostbind = _host.String() }
			if _name, err := sec.GetKey("name"); err == nil { option_hostname = _name.String() }
			if _aias, err := sec.GetKey("alias"); err == nil { option_hostalias = _aias.String() }
			if _dir, err := sec.GetKey("static_dir"); err == nil { option_staticdir = _dir.String() }
			if _url, err := sec.GetKey("static_url"); err == nil { option_staticurl = _url.String() }
			if _dir, err := sec.GetKey("templates"); err == nil { option_templatedir = _dir.String() }
		}
	}

	p.Set("server.host.bind", option_hostbind)
	p.Set("server.host.name", option_hostname)
	p.Set("server.host.alias", option_hostalias)
	p.Set("server.host.static_dir", option_staticdir)
	p.Set("server.host.static_url", option_staticurl)
	p.Set("server.host.template_dir", option_templatedir)
	p.Dump()
}

func main() {
	if os.Getenv("DEBUG") == "1" {
		p.Set("debug", "1")
	} else { p.Set("debug", "0") }

	setProps()
	router := s.Routes(p.Get("server.host.static_dir"), p.Get("server.host.static_url"))
	log.Fatal(http.ListenAndServe(p.Get("server.host.bind"), router))
}
