#!/bin/sh

# Install goose for migration
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run the migration
goose -dir database/migrations -table "migration_versions" postgres "$DATABASE_URL" up

# Start the API process
./main &
api=$!

# TODO: Add woeker later when work on scheduling job
# Start the worker process
# ./worker &
# worker=$!

# Wait for any processes to exit
wait $api # $worker

# Exit with status of process that exited first
exit $?
