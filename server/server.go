package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() {
	app := fiber.New()
	app.Listen(s.listenAddr)
}
