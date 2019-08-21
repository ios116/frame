package httpserver

import (
	"encoding/json"
	"github.com/ios116/frame/internal/accounts"
	"log"
	"net/http"
)

// GetUser get some accounts
func GetUser(w http.ResponseWriter, r *http.Request) {
	users := accounts.GetSomeUsers()
	b, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
