package fiber

import (
	"github.com/gofiber/fiber/v2"
	"gochat/cmd"
	"gochat/db/mongoDB"
	"gochat/types/constants"
	"log"
	"os"
)

type Server struct {
	AppName     string
	Version     string
	Port        string
	MongodbHost string
	MongodbPort string
	Database    string
	app         *fiber.App
}

func (s *Server) NewConfig(name string, version string) *Server {
	s.AppName = name
	s.Version = version
	s.setEnvValues()
	return s
}

func (s *Server) setEnvValues() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		s.Port = ":4080"
	} else {
		s.Port = ":" + port
	}

	mongoHost, ok := os.LookupEnv("MONGODB_HOST")

	if ok {
		s.MongodbHost = mongoHost
	} else {
		log.Fatal("MONGODB_HOST env variable not set")
	}

	mongoPort, ok := os.LookupEnv("MONGODB_PORT")
	if ok {
		s.MongodbPort = mongoPort
	}

	mongoDatabaseName, ok := os.LookupEnv("MONGODB_DATABASE")
	if ok {
		s.Database = mongoDatabaseName
	} else {
		log.Fatal("MONGODB_DATABASE env variable not set")
	}
}

func errorHandler(c *fiber.Ctx, err error) error {

	return c.Status(fiber.StatusBadRequest).JSON(constants.GlobalErrorHandlerResponse{
		Success: false,
		Message: err.Error(),
	})
}

func (s *Server) Run(preFork bool) {
	s.app = fiber.New(fiber.Config{AppName: s.AppName + " " + s.Version, CaseSensitive: true, Prefork: preFork, ErrorHandler: errorHandler})
	mongoDB.Client = mongoDB.Client.NewMongoClient(s.Database)
	mongoDB.Client.Connect(s.MongodbHost, s.MongodbPort)

	cmd.HandleMiddlewares(preFork, s.app)
	cmd.HandleRoutes(s.app)

	// Listener
	if err := s.app.Listen(s.Port); err != nil {
		log.Fatal("Error Occurred")
	}
}
