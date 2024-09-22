package server

import (
	"net/http"
)

type UserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (s *Server) users(w http.ResponseWriter, r *http.Request) {
	u, err := s.store.Users()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	users := []UserResponse{}
	for _, user := range u {

		users = append(users, UserResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (s *Server) user(w http.ResponseWriter, r *http.Request) {
	user, err := s.store.User(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, UserResponse{
		Id:    user.Id,
		Email: user.Email,
	})
}
