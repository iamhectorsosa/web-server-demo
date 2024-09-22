package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(map[string]string{"status": "OK"})
	if err != nil {
		log.Fatalf("error handling JSON marshal: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
