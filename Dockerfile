FROM golang:latest AS development
RUN apt-get update && apt-get install -y
WORKDIR /app
ENV APP_MODE=debug
COPY . /app
RUN go get -v
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]
EXPOSE 8080 3306

FROM golang:latest AS production
RUN apt-get update && apt-get install -y
WORKDIR /app
ENV APP_MODE=release
COPY . /app
RUN go build -o runner
EXPOSE 8080
ENTRYPOINT ["./runner"]