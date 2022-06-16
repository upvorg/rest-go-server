## Deploy Status

[![deploy](https://github.com/upvorg/server2/actions/workflows/deploy.yml/badge.svg)](https://github.com/upvorg/server2/actions/workflows/deploy.yml)

## Envrionment

- go version go1.18.1 darwin/amd64

## RUN

- go install github.com/cosmtrek/air@latest

```bash
# debug
env ENV=debug air
# sit
env ENV=release go run main.go

# release
cp ./.env /.env
go build -o /app
env ENV=release /app &
```

## Roadmap

- Collectmark

## Refs

- https://rapidapi.com/search/anime
- https://darjun.github.io/2020/04/04/godailylib/validator/
- https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
- https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

```bash
# tidy
go mod tidy

# run dev
env ENV=debug go run main.go

# run prod
env ENV=release go run main.go

# mysql dump
mysqldump --databases upv -u root -p > upv.sql
scp -rp root@upv.life:/~/upv.sql ./

# scp download files
scp -rp root@upv.life:/path/filename /path #将远程文件从服务器下载到本地

# scp upload files
scp -rp /Users/g/code/web/dist/admin/* root@upv.life:/var/www/admin
scp -rp /Users/g/code/web/dist/index/* root@upv.life:/var/www/html


# docker up (todo)
docker-compose up -d -f docker-compose.yml --env-file .env
docker-
```
