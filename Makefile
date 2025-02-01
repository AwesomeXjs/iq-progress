include .env.example
LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/main.go

up:
	docker-compose up -d

install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0


get-deps:
	go get github.com/labstack/echo/v4
	go get -u github.com/jackc/pgx/v4
	go get github.com/georgysavva/scany/pgxscan
	go get -u github.com/brianvoe/gofakeit
	go get -u github.com/Masterminds/squirrel
	go get github.com/joho/godotenv
	go get github.com/stretchr/testify/require
	go get github.com/golang/mock/gomock
	go get github.com/rs/cors
	go get go.uber.org/zap
	go get go.uber.org/zap/zapcore
	go get github.com/natefinch/lumberjack
	go get github.com/goccy/go-json

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

swagger:
	$(LOCAL_BIN)/swag init -g cmd/app/main.go

LOCAL_MIGRATION_DIR=$(CURDIR)/internal/migrations
LOCAL_MIGRATION_DSN="host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

init-migration:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} create first_tables sql