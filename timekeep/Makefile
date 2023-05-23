GOGLAGS=GOFLAGS=-mod=mod
GO=$(GOGLAGS) go

install: build
	@sudo mv tk /usr/bin

build:
	@$(GO) build -o tk

test:
	@$(GO) test -v -count=1 ./... && go mod tidy
