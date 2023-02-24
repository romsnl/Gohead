FROM golang:1.20-alpine AS buildo

WORKDIR /build

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/gohead

FROM scratch

WORKDIR /

COPY --from=buildo /build/bin/gohead /gohead

EXPOSE 8080

CMD [ "./gohead" ]