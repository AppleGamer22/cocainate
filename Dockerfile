FROM golang:1.17.8-alpine3.15 AS build
WORKDIR /cocainate
ENV CGO_ENABLED 0
COPY . .
RUN go build .

FROM ubuntu:20.04
RUN sudo apt install -y ubuntu-desktop-minimal