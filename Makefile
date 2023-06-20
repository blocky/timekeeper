GOGLAGS=GOFLAGS=-mod=mod
GO=$(GOGLAGS) go

GIT_COMMIT=$(shell git rev-parse --short HEAD)
CONFIGS=cmd/configs
TASKS_PHASE_2=$(CONFIGS)/tasks-phase-2.json
TASKS_MEETINGS=$(CONFIGS)/tasks-meetings.json
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
		--project 5f91ec0fb1d41c38c2d6719b --json > $(TASKS_PHASE_2)
	@clockify-cli tasks list \
		--project 5f47d5879d6dc04fbfedcdab --json > $(TASKS_MEETINGS)
	@jq -s 'add' $(CONFIGS)/*.json \
		| jq 'map(. | select(.status=="ACTIVE") | {id, name, projectId})' \
		| jq -s '.[] | sort_by(.name)' > $(TASKS)