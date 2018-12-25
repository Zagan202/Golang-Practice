# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=cryptobin

all: run print clean
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME) *.enc
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
print:
		@echo "-------------------------------------"