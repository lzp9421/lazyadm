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
	host := conf.StringDefault("http.host", "127.0.0.1")
	port := conf.IntDefault("http.port", 8080)
	conf.SetSection("DEFAULT")
	return &Server{
		host:   host,
		port:   port,
		router: router,
	}, nil
}

func (this *Server) Listen() {
	fmt.Print(this.host + ":" + strconv.Itoa(this.port))
	http.ListenAndServe(this.host+":"+strconv.Itoa(this.port), this.router.GetHandle())
}
