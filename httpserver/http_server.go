package httpserver

import (
	"fmt"
	"githab.com/gorrila/mux"
	"net/http"
)

type HttpServer struct {
	httpPort int
	router   http.Handler
}

func (s *HttpServer) Start() error {
	return http.ListenAndServer(fmt.Sprintf("0.0.0.0:%d", s.httpPort), s.router)
}

func NewHTTPServer(port int) *HttpServer {

	r := mux.NewRouter()
	hs = HttpServer{router: r, httpPort: port}

	http.Handle("/", r)

	return &hs

}
