package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vicolby/meets/storer"
	"github.com/vicolby/meets/types"
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
	log.Println("Server started on", s.listenAddr)
}

func (s *Server) handleGetUserByID(c *fiber.Ctx, id int) (*types.User, error) {
	user, err := s.userStore.GetUser(id)
	if err != nil {
		return &types.User{}, err
	}
	return user, nil
}

func (s *Server) handleGetEventByID(c *fiber.Ctx) (*types.Event, error) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &types.Event{}, err
	}
	event, err := s.eventStore.GetEvent(id)
	if err != nil {
		return &types.Event{}, err
	}
	return event, nil
}

func (s *Server) handleGetEvents(c *fiber.Ctx) ([]*types.Event, error) {
	events, err := s.eventStore.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Server) handleAddEvent(c *fiber.Ctx) error {
	event := &types.Event{}
	if err := c.BodyParser(event); err != nil {
		return err
	}
	if err := s.eventStore.AddEvent(event); err != nil {
		return err
	}
	return c.JSON(event)
}
