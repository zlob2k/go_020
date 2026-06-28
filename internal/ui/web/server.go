package web

// internal/ui/web

import (
	"fmt"
	"net/http"
)

type Server struct {
	ServerAddr string
	Router     http.Handler
}

func NewServer(server_addr string, router http.Handler) *Server {
	fmt.Printf("\n NewServer()")
	return &Server{
		ServerAddr: server_addr,
		Router:     router,
	}
}

func (s *Server) Start() error {
	fmt.Printf("\n Start()")
	fmt.Printf("\n ServerAddr = %s", s.ServerAddr)
	return http.ListenAndServe(s.ServerAddr, s.Router)
}
