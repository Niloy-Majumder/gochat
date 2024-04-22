package fiber

import (
	"github.com/gofiber/fiber/v2"
	"gochat/cmd"
	"log"
)

type Server struct {
	AppName string
	Version string
	Port    string
	app     *fiber.App
}

func (s *Server) NewDevConfig(name string, version string) *Server {
	s.AppName = name
	s.Version = version
	s.Port = ":4080"
	return s
}

func (s *Server) NewProdConfig(name string, version string) *Server {
	s.AppName = name
	s.Version = version
	s.Port = ":8080"
	return s
}

func (s *Server) Serve(preFork bool) {
	s.app = fiber.New(fiber.Config{AppName: s.AppName + " " + s.Version, CaseSensitive: true, Prefork: preFork})

	cmd.HandleMiddlewares(preFork, s.app)
	cmd.HandleRoutes(s.app)

	// Listener
	if err := s.app.Listen(s.Port); err != nil {
		log.Fatal("Error Occurred")
	}
}
