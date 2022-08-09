all: build

gocyclo:
	gocyclo -over 15 .

golangci-lint:
	golangci-lint run ./...

staticcheck:
	staticcheck -checks 'all,-ST*' ./...

tidy:
	go mod tidy

fmt: tidy
	go fmt ./...

lint: fmt staticcheck golangci-lint gocyclo

test: lint
	go test ./... -cover

run:
	go run ./cmd/go-photos/main.go

debug:
	go run ./cmd/go-photos/main.go --debug

build: lint test
	go build -ldflags "-s -w" -o ./bin/go-photos ./cmd/go-photos/main.go

.NOTPARALLEL:

.PHONY: all gocyclo golangci-lint staticcheck tidy fmt lint test run debug build
