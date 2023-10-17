FROM ubuntu:latest

COPY ./bin /places
COPY ./configs /places

CMD chmod +x /places/suggest

WORKDIR /places
