package server

import (
	"blog/api/src/infrastructure/routes"
	"net/http"
)


func NewServer(port string) *http.Server {
	r := routes.NewRouter();
	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	return server
}