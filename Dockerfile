FROM golang:1.20 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-s" -o=./bin/cardamom ./cmd/web

FROM gcr.io/distroless/static-debian11

LABEL org.opencontainers.image.source=https://git.sr.ht/~edwlarkey/cardamom
LABEL org.opencontainers.image.description="Webapp to store bookmarks/articles to read later "
LABEL org.opencontainers.image.licenses=AGPL-3.0-or-later

COPY --from=build /go/src/app/bin/cardamom /cardamom
CMD ["/cardamom", "--dsn", "/app/data/cardamom.sqlite?_time_format=sqlite"]
