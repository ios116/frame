package httpserver

import (
	"github.com/gorilla/mux"
)

func (s *HTTPServer) createRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	subRouterAPI := router.PathPrefix("/api").Subrouter()
	subRouterAPI.HandleFunc("/user", s.GetUser).Methods("GET")
	//router.NotFoundHandler = http.HandlerFunc(controllers.Page404)
	return router
}
