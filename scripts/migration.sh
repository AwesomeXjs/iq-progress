#!/bin/bash
source .env.example

sleep 2 && goose -dir "migrations" postgres "${MIGRATION_DSN}" up -v