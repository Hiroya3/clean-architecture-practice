PROJECTNAME=todo-app

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

init:
	make build-debug

build:
	GOOS=linux CGO_ENABLED=0 go build -o $(GOBIN)/$(PROJECTNAME) ./cmd

.PHONY: build-debug
build-debug:
	go mod download
	type dlv || go install github.com/go-delve/delve/cmd/dlv@v1.8.0
	mkdir -p tmp
	go build -gcflags "all=-N -l" -o ./tmp/todo-app ./cmd