default: help

ifeq ($(OS),Windows_NT)
    RM = rmdir /S /Q
else
    RM = rm -rf
endif

.PHONY: compose.up
compose.up:
	@which docker
	@docker compose -f ./dev/docker-compose.yml up -d
	@docker compose -f ./dev/docker-compose.yml exec postgres \
		sh -c 'while ! nc -z postgres 5432; do sleep 0.1; done'

.PHONY: compose.down
compose.down:
	@which docker
	@docker compose -f dev/docker-compose.yml down

.PHONY: migrate.local.up
migrate.local.up: compose.up
	@which go
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0\
		-path=./dev/migrate/postgres -database 'postgres://krivenkov@localhost:5432/order?sslmode=disable' up

.PHONY: migrate.local.down
migrate.local.down: compose.up
	@which go
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0\
    	-path=./dev/migrate/postgres -database 'postgres://krivenkov@localhost:5432/order?sslmode=disable' down

.PHONY: lint
## Lint files
lint:
	@which go
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1\
		run -v -c=.golangci.yml ./...

.PHONY: test
test:
	@which go
	@go test -v -cover -gcflags=-l ./internal/...

run:
	go run ./cmd/order-api/main.go --cfg=./res/cfg-local.yml

generate-proto:
	bash scripts/compile-proto.sh

## Generates http server from swagger
gen.server:
	$(RM) $(call FIXPATH,./internal/server/http/models) &&\
	$(RM) $(call FIXPATH,./internal/server/http/operations) &&\
	go run github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0\
    generate server -f ./api-spec/swagger.json\
    --exclude-main\
    --server-package=./internal/server/http\
    --model-package=./internal/server/http/models &&\
  git add ./internal/server/http

.PHONY: generate
generate:
	@which go
	@go generate -x ./...