package main

import (
	"log"

	"github.com/iamhectorsosa/web-server-demo/internal/memorystore"
	"github.com/iamhectorsosa/web-server-demo/internal/server"
)

func main() {
	store := memorystore.New()
	server := server.New(store)

	log.Printf("Listening on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
