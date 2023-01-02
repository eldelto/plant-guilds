.DELETE_ON_ERROR:

# High-level commands
.PHONY: build
build: bin/plant-guilds

.PHONY: init
init: download
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: download
download:
	@go mod download
	@npm install --prefix analyzer-web/

.PHONY: run
run: build
	@echo Starting Plant-Guilds ...
	@bin/plant-guilds

.PHONY: run-loop
run-loop:
	@reflex -r "\.(go|tmpl|css|js)$$" -R ".*node_modules.*" -s make run

.PHONY: test
test: lint
	@echo Testing ...
	@go test ./...

.PHONY: lint
lint:
	@go mod tidy
	@go fmt ./...
	@staticcheck ./...

.PHONY: test-loop
test-loop:
	@echo Waiting for file changes ...
	@reflex -r "\.(go|html|css|js)$$" -R ".*node_modules.*" make test

.PHONY: clean
clean:
	@rm -f bin/*
	@rm -rf flight-controller/build

# Plant-Guilds
bin/plant-guilds: .FORCE
	@go build -o bin/plant-guilds cmd/plant-guilds/main.go

.FORCE:
