all: tidy build run

init:
	go mod download github.com/gorilla/mux
	go mod download github.com/spf13/cobra
	go mod download github.com/vanng822/go-solr
	go fmt
	go mod tidy
	go mod vendor
	mkdir bin
	build
	run

tidy:
	go mod tidy
	go fmt ./...

build:
	go test ./...
	go build -o bin/sopagex main.go

run:
	chmod +x bin/sopagex
	chmod +x sopagex.sh
	./sopagex.sh

install:
	go install -v ./...