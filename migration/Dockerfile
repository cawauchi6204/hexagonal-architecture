# golang-migrateを使用してマイグレーションを行う
FROM golang:1.22.4-alpine

ENV ROOT=/go/src/app/migration
ENV CGO_ENABLED=0
WORKDIR ${ROOT}

RUN apk update --no-cache && \
    apk add --no-cache --update git \
    mysql-client

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

COPY . .
COPY run_migration.sh /run_migration.sh
RUN chmod +x /run_migration.sh

CMD ["/run_migration.sh"]