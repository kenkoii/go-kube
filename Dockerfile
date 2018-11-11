FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp

EXPOSE 3000
ENTRYPOINT ./goapp