package httpserver

import (
	"encoding/json"
	"github.com/ios116/frame/internal"
	"log"
	"net/http"
)

// GetUser get some users
func GetUser(w http.ResponseWriter, r *http.Request) {
	users := intetrnal.GetSomeUsers()
	b, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
