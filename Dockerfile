FROM ubuntu:latest

COPY ./bin /places
COPY ./configs /places

WORKDIR /places
