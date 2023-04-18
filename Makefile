.PHONY: build
build:
	@go build ./cmd/api
clean:
	@rm ./api

.DEFAULT_GOAL := build