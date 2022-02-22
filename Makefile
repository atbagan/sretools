.PHONY: all
build: test build

.PHONY: deps
deps:
	go get -v ./...

.PHONY: test
test:
	go get -u golang.org/x/lint/golint
	$(GOPATH)/bin/golint ./...
	go test ./...

.PHONY: build
build:
	golint ./...
	go test ./...
	go build -o /usr/local/bin/sretools