FROM golang:1.21.12-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update --no-cache && \
    apk add --no-cache --update git \
    mysql-client

RUN go install github.com/cosmtrek/air@v1.43.0

COPY go.mod go.sum ./
# COPY . ./
RUN go mod download
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]