FROM ubuntu:latest

COPY ./bin/ /places/
COPY ./config.toml /places/

WORKDIR /places
