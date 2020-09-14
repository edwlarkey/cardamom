APP=cardamom
PROJECT_PATH=$(realpath $(shell dirname $(lastword $(MAKEFILE_LIST))))


build: generate
	go build -o ${APP} cmd/web/*

run:
	go run -race main.go

test:
	go test cmd/web/*

generate:
	go generate pkg/assets/assets.go

lint:
	docker run --rm -v ${PROJECT_PATH}:/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint run -v

clean:
	rm -f cmd/web/rice-box.go

.PHONY: build run test generate clean lint
