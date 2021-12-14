# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -p 1
GOGET=$(GOCMD) get
GOVENDOR=$(GOCMD) mod vendor
BINARY_NAME=go-imx-client
MAIN_FILE=main.go

all: test build
build: 
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
test: 
	$(GOTEST)  ./...
clean: 
	$(GOCLEAN)
run: 
	$(GORUN) $(MAIN_FILE)
deps: 
	$(GOGET) $(COMMON)
	$(GOVENDOR)