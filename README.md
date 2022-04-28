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

## Refs

- https://rapidapi.com/search/anime
