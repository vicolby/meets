package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vicolby/meets/storer"
)

type Server struct {
	listenAddr string
	userStore  storer.UserStorage
	eventStore storer.EventStorage
}

func NewServer(listenAddr string, userStore storer.UserStorage, eventStore storer.EventStorage) *Server {
	return &Server{
		listenAddr: listenAddr,
		userStore:  userStore,
		eventStore: eventStore,
	}
}

func (s *Server) Start() {
	app := fiber.New()
	app.Listen(s.listenAddr)
}
