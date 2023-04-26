FROM golang:1.20.3-alpine3.17

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum main.go ./
RUN go mod download && go mod verify && go build -v -o /usr/local/bin/app ./...

CMD ["app"]