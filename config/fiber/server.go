package fiber

import (
	"errors"
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
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(constants.GlobalErrorHandlerResponse{
		Success: false,
		Message: e.Message,
	})
}
func _initMongoIndexes(client *mongoDB.MongoClient) {
	err := client.SetIndexes("users", "email")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Run(preFork bool) {
	s.app = fiber.New(fiber.Config{AppName: s.AppName + " " + s.Version, CaseSensitive: true, Prefork: preFork, ErrorHandler: errorHandler})

	// MongoDB Operations -
	mongoDB.Client = mongoDB.Client.NewMongoClient(s.Database)
	mongoDB.Client.Connect(s.MongodbHost, s.MongodbPort)
	_initMongoIndexes(mongoDB.Client)

	cmd.HandleMiddlewares(preFork, s.app)
	cmd.HandleRoutes(s.app)

	// Listener
	if err := s.app.Listen(s.Port); err != nil {
		log.Fatal("Error Occurred")
	}
}
