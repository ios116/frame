package httpserver

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// HTTPServer struct for init http server
type HTTPServer struct {
	Port   int
	Host   string
	Logger *zap.Logger
}

// Start start http server
func (s *HTTPServer) Start() {
	r := s.createRouters()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler: r,
	}
	s.Logger.Info("Server is starting", zap.String("host", s.Host), zap.Int("port", s.Port))
	s.Logger.Fatal(server.ListenAndServe().Error())
}
