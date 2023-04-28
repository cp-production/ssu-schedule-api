.PHONY: build
build:
	@go build -v ./cmd/api
clean:
	@rm ./api

.DEFAULT_GOAL := build