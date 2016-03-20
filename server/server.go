package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NeoHuang/go-webhooks/handler"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	log.Printf("Webhook Server Starting on PORT: %d\n", server.config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", server.config.Port), nil); err != nil {
		log.Panicf("WEBHOOK SERVER SHUTTING DOWN (%s)\n\n", err)
	}
}

func AddHandler(pattern string, eventHandler handler.Handler) {
	log.Printf("Setting up route... %s", pattern)
	http.Handle(pattern, NewHttpHandler(eventHandler))
}
