package http

import (
	"net/http"

	"github.com/weriKK/microservice"
)

type httpServer struct {
	server *http.Server
}

func NewServer(addr string, handler http.Handler) microservice.HTTPAPIServer {

	server := httpServer{
		server: &http.Server{
			Addr:      addr,
			TLSConfig: nil,
			Handler:   handler,
		},
	}

	return &server
}

func (s *httpServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	return s.Shutdown()
}
