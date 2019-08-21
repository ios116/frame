package httpserver

import (
	"github.com/gorilla/mux"
)

func createRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	subRouterApi := router.PathPrefix("api").Subrouter()
	subRouterApi.HandleFunc("/user", GetUser).Methods("GET")
	//router.NotFoundHandler = http.HandlerFunc(controllers.Page404)
	return router
}
