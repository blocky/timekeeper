GOGLAGS=GOFLAGS=-mod=mod
GO=$(GOGLAGS) go

GIT_COMMIT=$(shell git rev-parse --short HEAD)
CONFIGS=cmd/configs
EPICS=$(CONFIGS)/epics.json
MEETINGS=$(CONFIGS)/meetings.json
TIMEOFF=$(CONFIGS)/timeoff.json
TASKS=$(CONFIGS)/tasks.json

configs:
	@mkdir -p $(CONFIGS)

install: build
	@sudo mv tk /usr/bin

build: tasks
	@$(GO) build -ldflags "-X main.Version=$(GIT_COMMIT)" -o tk

test:
	@$(GO) test -v -count=1 ./... && go mod tidy

tasks: configs
	@clockify-cli tasks list \
		--project 5f91ec0fb1d41c38c2d6719b --json > $(EPICS)
	@clockify-cli tasks list \
		--project 5f47d5879d6dc04fbfedcdab --json > $(MEETINGS)
	@clockify-cli tasks list \
		--project 61a4d4563b4281137e936e9d --json > $(TIMEOFF)
	@jq -s 'add' $(CONFIGS)/*.json \
		| jq 'map(. | select(.status=="ACTIVE") | {id, name, projectId})' \
		| jq -s '.[] | sort_by(.name)' > $(TASKS)