FROM golang:1.20 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-s" -o=./bin/cardamom ./cmd/web

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/src/app/bin/cardamom /cardamom
CMD ["/cardamom", "--dsn", "/app/data/cardamom.sqlite?_time_format=sqlite"]
