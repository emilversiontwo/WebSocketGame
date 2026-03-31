FROM golang:1.26.1
LABEL authors="soupik"

WORKDIR /usr/src/app

COPY ./server/go.mod ./go.mod

RUN go mod download

COPY ./server .