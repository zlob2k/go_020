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
	fmt.Printf("\nNewServer()")
	return &Server{
		ServerAddr: server_addr,
		Router:     router,
	}
}

func (s *Server) Start() error {
	fmt.Printf("\nStart() ServerAddr = %s", s.ServerAddr)
	err := http.ListenAndServe(s.ServerAddr, s.Router)
	return err
}
