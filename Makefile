.PHONY: lint test run generate-proto generate-api

lint:
	sh ./scripts/lint.sh

test:
	sh ./scripts/test.sh

run:
	go run ./cmd/order-api/main.go --cfg=./res/cfg-local.yml

generate-proto:
	bash scripts/compile-proto.sh

generate-api:
	bash scripts/gensrv.sh
