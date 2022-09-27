FROM golang:alpine AS development
RUN apt-get update && apt-get install -y
WORKDIR /app
ENV ENV=debug
COPY . /app
RUN go get -v
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]
EXPOSE 8080 3306

FROM alpine:latest AS production
RUN apt-get update && apt-get install -y
WORKDIR /app
ENV ENV=release
RUN go build -o app
EXPOSE 8080
ENTRYPOINT ["./app"]