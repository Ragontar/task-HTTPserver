package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*

	Server-related stuff:

	ConfigData struct --- consists of Server and Database config structures. Used to parse a json config.

	Server struct --- consists of server parameters. Also sed to parse a json config.

	getConfig() *ConfigData --- parses a json config file and returns a filled ConfigData. Panics if
		is not able to read a config file.

	NewServer() *Server --- initializes a new HTTP server, returns Server structure. Uses getConfig().

	(s *Server) Run() --- starts the server.

*/

type ConfigData struct{
	Serv Server `json:"server"`
	Db Database `json:"database"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func getConfig() *ConfigData {
	var conf ConfigData

	configPath := "config.json"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &conf)

	return &conf
}

func NewServer() *Server {
	conf := getConfig()
	return &Server{conf.Serv.Host, conf.Serv.Port}
}

func (s *Server) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			fmt.Fprintln(w, "What's wrong with you?")
			return
		}
		if r.Method == http.MethodGet {
			usr, err := getUser(name)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			fmt.Fprintf(w, "User %s found in database! Take him!\n", usr.Name)
		}
		if r.Method == http.MethodPost {
			err := addUser(User{name})
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			fmt.Fprintf(w, "User %s added to database!\n", name)
		}
	})

	addr := ":" + s.Port
	http.ListenAndServe(addr, nil)
}