APP=cardamom


build: generate
	go build -o ${APP} cmd/web/*

run:
	go run -race main.go

test:
	go test cmd/web/*

generate:
	go generate pkg/assets/assets_generate.go

clean:
	rm -f cmd/web/rice-box.go

.PHONY: build run test generate clean
