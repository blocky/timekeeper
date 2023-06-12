GOGLAGS=GOFLAGS=-mod=mod
GO=$(GOGLAGS) go

GIT_COMMIT=$(shell git rev-parse --short HEAD)

install: build
	@sudo mv tk /usr/bin

build:
	@$(GO) build -ldflags "-X main.Version=$(GIT_COMMIT)" -o tk

test:
	@$(GO) test -v -count=1 ./... && go mod tidy
