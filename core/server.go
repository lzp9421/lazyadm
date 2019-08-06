package core

import (
	"github.com/revel/config"
	"fmt"
	"strconv"
	"net/http"
)

type Server struct {
	host   string
	port   int
	router *Router
}

func NewServer(conf *config.Context, router *Router) (*Server, error) {
	conf.SetSection("server")
	defer conf.SetSection("DEFAULT")

	host := conf.StringDefault("http.host", "127.0.0.1")
	port := conf.IntDefault("http.port", 8080)

	return &Server{
		host:   host,
		port:   port,
		router: router,
	}, nil
}

func (s *Server) Listen() {
	fmt.Print(s.host + ":" + strconv.Itoa(s.port))
	http.ListenAndServe(s.host+":"+strconv.Itoa(s.port), s.router.GetHandle())
}
