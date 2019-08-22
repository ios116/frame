package httpserver

import (
	"encoding/json"
	"github.com/ios116/frame/internal/accounts"
	"net/http"
)

// GetUser get some accounts
func (s *HTTPServer) GetUser(w http.ResponseWriter, r *http.Request) {
	users := accounts.GetSomeUsers()
	b, err := json.Marshal(users)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
