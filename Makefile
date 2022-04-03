PROJECTNAME=todo-app

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

build:
	GOOS=linux CGO_ENABLED=0 go build -o $(GOBIN)/$(PROJECTNAME) ./cmd