# Order

This repository contains service designed to prepare order's data.
And also it provides grpc/http api for order.

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

## Database migration
resides in `scripts/migrate/postgres`

## Elastic mappings
resides in `scripts/migrate/elastic`

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
