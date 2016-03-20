package main

import (
	"github.com/NeoHuang/go-webhooks/handler"
	"github.com/NeoHuang/go-webhooks/server"
)

func main() {
	config := server.NewConfig("config.json")
	httpServer := server.NewServer(config)
	server.AddHandler("/log", handler.NewLogHandler())
	httpServer.Start()
}
