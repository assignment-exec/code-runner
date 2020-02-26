# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY=code-runner-server

all: build
build:
	$(GOBUILD) -o $(BINARY)

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

run: build
	./$(BINARY)

dependency:
	$(GOMOD) download