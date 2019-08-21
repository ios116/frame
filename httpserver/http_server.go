package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

// HttpServer struct for init http server
type HttpServer struct {
	Port int
	Host string
}

// Start start http server
func (s *HttpServer) Start() {
	r := createRouters()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
