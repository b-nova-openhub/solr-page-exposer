all: test vet fmt mod build run

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

mod:
	go mod tidy
	go mod vendor

build:
	go build -o bin/sopagex cmd/sopagex/main.go

run:
	chmod +x sopagex.sh
	./sopagex.sh

install:
	mod
	go install -v ./...
