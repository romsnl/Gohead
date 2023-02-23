# Gohead

Simple golang webapp serving http.Header

## Getting Started

```sh
❯ curl 127.0.0.1:8080
Method: GET
Host: 127.0.0.1:8080
Remote Address: 127.0.0.1:54013
User Agent: curl/7.86.0
Request URI: /
Headers:
Accept: */*
User-Agent: curl/7.86.0
```

### Build & Run

> go bin

```sh
go build -o bin/gohead
./bin/gohead
```

> Dockerfile

```sh
docker build -t gohead .
docker run -it -p 8080:8080 gohead
```
