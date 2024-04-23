package main

import (
	"flag"
	"github.com/joho/godotenv"
	"gochat/config/fiber"
	"log"
)

var (
	prod = flag.Bool("prod", false, "Enable Production Server")
)

func main() {
	flag.Parse()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	server := &fiber.Server{}
	server = server.NewConfig("Go Chat Application", "1.0.0")

	server.Run(*prod)
}
