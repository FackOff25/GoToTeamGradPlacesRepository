FROM ubuntu:latest

COPY ./bin/ /places/
COPY ./configs/config.toml /places/

WORKDIR /places
