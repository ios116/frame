package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	Port int
	Host string
	Router   http.Handler
}

func (s *HttpServer) Start() {
	r := createRouters()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
