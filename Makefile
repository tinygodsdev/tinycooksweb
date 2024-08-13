ifneq ("$(wildcard env/.env.local)","")
    include env/.env.local
    export $(shell sed 's/=.*//' env/.env.local)
endif

.PHONY: run
run: entgen css
	go run cmd/main.go

.PHONY: seed
seed: entgen
	go run cmd/tools/seed/main.go

.PHONY: build
build:
	go build -o bin/main cmd/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: deploy
deploy:
	git push dokku main

.PHONY: tidy
tidy:
	@echo "Setting up Git for token-based authentication"
	@git config --global url."https://oauth2:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
	@echo "Setting GOPRIVATE environment variable"
	@GOPRIVATE=github.com/tinygodsdev/*
	@echo "Running go mod tidy"
	@go mod tidy

.PHONY: deps
deps:
	go get github.com/tinygodsdev/datasdk
	go mod tidy

.PHONY: entgen
entgen:
	cd ./pkg/storage/entstorage && go generate ./ent

.PHONY: css
css:
	cd web && npm run build

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
	./bin/golangci-lint run -v
