FROM ubuntu:latest

COPY ./bin/ /places/
COPY ./configs/config.yaml /places/

WORKDIR /places
