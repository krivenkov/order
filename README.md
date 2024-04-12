# Order

This repository contains service designed to prepare order's data.
And also it provides grpc/http api for order.

## Listen to events
- User update (user.update.user.1)

## Interfaces

### GRPC
#### Inner
- [order](pkg/proto/api/service.proto)

### HTTP
- [order api](api-spec/swagger.json)

## External dependencies
- Postgres
- ElasticSearch
- Keycloak
- Kafka

## Database migration
resides in `dev/migrate/postgres`

## Elastic mappings
resides in `dev/migrate/elastic`

## DB migrations
You can start pg migration with command
```
$ make migrate.local.up
```

## Tests
You can start tests with command
```
$ make test
```

## Lints
You can start tests with command
```
$ make lint
```

## Generate proto
You can generate proto with command
```
$ make generate-proto
```

## Generate api
You can generate api with command
```
$ make generate-api
```

## Dependencies
```
$ go mod tidy
```

## Format
#### JetBrains
- Enable checkbox "Reformat code" in commit dialog
- Enable checkbox "Optimize imports" in commit dialog
#### Other IDE
Use gofmt, then goimports

## Run
#### Local environment
```
$ make run
```
