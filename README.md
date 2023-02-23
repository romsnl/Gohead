# Gohead

Simple golang webapp serving http.Header

## Getting Started

```sh
❯ curl 127.0.0.1:8080
Host: 127.0.0.1:8080
Method: GET
URL: /
Proto: HTTP/1.1
ContentLength: 0
TransferEncoding: []
Close: false
Trailer: map[]
RemoteAddr: 172.17.0.1:57230
User Agent: curl/7.86.0
Remote Address: 172.17.0.1:57230

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
