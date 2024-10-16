#!/bin/sh

if [ -z "$MIGRATION_DIRECTION" ]; then
    echo "Error: MIGRATION_DIRECTION is not set. Use 'up' or 'down'."
    exit 1
fi

migrate -path ./migrations -database "mysql://root:rootpassword@tcp(db:3306)/db" $MIGRATION_DIRECTION