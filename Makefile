ADDR ?= :8090
BASE_URL ?= http://localhost:8090
GOBIN := $(shell go env GOPATH)/bin
TEMPL ?= $(GOBIN)/templ
AIR ?= $(GOBIN)/air
TAILWINDCSS ?= npx @tailwindcss/cli

.PHONY: help install-tools templ css css-watch test build run dev clean

help:
	@echo "Targets:"
	@echo "  make install-tools  Install Go dev tools: templ and air"
	@echo "  make templ          Generate templ Go files"
	@echo "  make css            Build assets/css/app.css from Tailwind input"
	@echo "  make test           Generate templates and run Go tests"
	@echo "  make build          Generate templates and build ./tmp/server"
	@echo "  make run            Run the Go server"
	@echo "  make dev            Run Air live reload if installed, otherwise run normally"
	@echo "  make clean          Remove local build output"

install-tools:
	go install github.com/a-h/templ/cmd/templ@v0.3.1020
	go install github.com/air-verse/air@latest
	npm install -D tailwindcss @tailwindcss/cli

templ:
	@test -x "$(TEMPL)" || (echo "templ not found at $(TEMPL). Run 'make install-tools'." && exit 1)
	$(TEMPL) generate

css:
	$(TAILWINDCSS) -c tailwind.config.js -i assets/css/input.css -o assets/css/app.css

css-watch:
	$(TAILWINDCSS) -c tailwind.config.js -i assets/css/input.css -o assets/css/app.css --watch

test: templ
	go test ./...

build: templ
	mkdir -p tmp
	go build -o ./tmp/server ./cmd

run: templ
	ADDR=$(ADDR) BASE_URL=$(BASE_URL) go run ./cmd

dev:
	@if test -x "$(AIR)"; then \
		ADDR=$(ADDR) BASE_URL=$(BASE_URL) $(AIR); \
	else \
		echo "air not found at $(AIR). Falling back to make run."; \
		$(MAKE) run; \
	fi

clean:
	rm -rf tmp
