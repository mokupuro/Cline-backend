FROM golang:latest

RUN mkdir /app
WORKDIR /app

ENTRYPOINT ["./entrypoint.sh"]