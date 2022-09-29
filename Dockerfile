FROM golang:latest AS development
RUN apt-get update && apt-get install -y
WORKDIR /app
ENV APP_MODE=debug
COPY . /app
RUN go get -v
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]
EXPOSE 8080

## 打包取消注释
# FROM golang:latest AS builder
# WORKDIR /app
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -o runner -a -installsuffix .

# FROM alpine:latest AS production
# WORKDIR /app
# ENV APP_MODE=release
# COPY --from=builder /app/runner /app/runner
# COPY --from=builder /app/.env /app/.env
# EXPOSE 8080
# ENTRYPOINT ["./runner"]