build:
	# Build The Application
	go build -o bin/main .

compile:
	# Build For All Operating Systems
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

dev_requirements:
	# Dev Dependencies
	go install github.com/cosmtrek/air@latest
	air init

requirements:
	# Dependencies
	go mod tidy

start_dev:
	# Running The Application With Live Reloading
	air -build.bin "./bin/main" -build.cmd "go build -o ./bin/main ."

start_prod: build
	# Start The Application
	./bin/main -prod

all: compile start

