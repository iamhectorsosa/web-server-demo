package server

import (
	"log"
	"net/http"
	"os"

	"github.com/iamhectorsosa/web-server-demo/internal/store"
	"github.com/joho/godotenv"
)

type Server struct {
	store store.Store
	http.Server
}

func New(store store.Store) *Server {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := ":" + os.Getenv("PORT")

	server := new(Server)
	server.store = store

	router := http.NewServeMux()
	router.HandleFunc("GET /api/health", server.health)
	router.HandleFunc("GET /api/users", server.users)
	router.HandleFunc("GET /api/users/{id}", server.user)

	server.Addr = addr
	server.Handler = router

	return server
}
