app_name=cardamom

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## watch: run air to lbuild and live reload when files are changed
.PHONY: watch
watch:
	air

## install/tools: installs development tools
.PHONY: install/tools
install/tools:
	@echo 'Installing migrate tool'
	go install -tags 'sqlite' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database sqlite://$(app_name).sqlite up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
#
## test: run tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -race -v -fullpath ./...

## cover: run tests and show coverage report
.PHONY: cover
cover:
	@echo 'Running tests...'
	go test -coverprofile=coverage.out -race -v -fullpath ./...
	go tool cover -html=coverage.out

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -v -fullpath -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build: build the cmd/web application
.PHONY: build
build:
	@echo 'Building cmd/web...'
	go build -ldflags="-s" -o=./bin/$(app_name) ./cmd/web

## build/docker: build the cmd/web application docker image
.PHONY: build/docker
build/docker:
	@echo 'Building docker image...'
	docker build . -t cardamom:latest
	docker build . -t ghcr.io/edwlarkey/cardamom:latest

## release: build and push docker image
.PHONY: release
release: build/docker
	@echo 'Pushing docker image...'
	docker push ghcr.io/edwlarkey/cardamom:latest
