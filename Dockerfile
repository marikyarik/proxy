# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go install github.com/valyala/quicktemplate/qtc

RUN qtc && go build -o /proxy

EXPOSE 80

CMD [ "/proxy" ]