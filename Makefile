.DEFAULT_GOAL := build

## Only Format
fmt:
	go fmt ./...
.PHONY:fmt

## Format and Lint
lint: fmt
	golint ./...
.PHONY:lint

## Format and Vet
vet: fmt
	go vet ./...
.PHONY:vet

## Build the app
build: vet
	go build 
.PHONY:build
