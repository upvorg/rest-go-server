## Deploy Status

[![deploy](https://github.com/upvorg/server2/actions/workflows/deploy.yml/badge.svg)](https://github.com/upvorg/server2/actions/workflows/deploy.yml)

## Envrionment

- go version go1.18.1 darwin/amd64

```bash
# debug
env ENV=debug fresh
# sit
env ENV=release go run main.go

# release
cp ./.env /.env
go build -o /app
env ENV=release /app &
```

## Go Packages

- go get -d github.com/pilu/fresh

## Roadmap

- Collectmark

## Refs

- https://rapidapi.com/search/anime
- https://darjun.github.io/2020/04/04/godailylib/validator/
