package main

import (
	"flag"
	"gochat/config/fiber"
)

var (
	prod = flag.Bool("prod", false, "Enable Production Server")
)

func main() {
	flag.Parse()
	server := &fiber.Server{}
	if *prod {
		server = server.NewProdConfig("Go Chat Application", "1.0.0")
	} else {
		server = server.NewDevConfig("Go Chat Application", "1.0.0")
	}
	server.Serve(*prod)
}
