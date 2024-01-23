FROM ubuntu:latest
LABEL authors="user"

ENTRYPOINT ["top", "-b"]

FROM golang:1.21.6-alpine

WORKDIR