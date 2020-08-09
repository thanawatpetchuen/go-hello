FROM golang:1.12-alpine

WORKDIR /app/echo

COPY . .

RUN apk add --no-cache git

RUN go get -u github.com/cosmtrek/air

RUN go mod vendor

CMD air