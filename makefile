BINARY_NAME := vending_machine
GO := go

.phony: build

build: 
	$(GO) build -v -o $(BINARY_NAME) .