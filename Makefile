-include .env
export

##${APP_LISTENER}

SERVICE_NAME ?= $(shell echo ${APP_NAME} | sed 's/-/_/g')

.PHONY: run
run:
	go run .\cmd\vacancy_service\.

.PHONY: test
test:
	go test -v ./...

.PHONY: test-unit
test-unit:
	go test -tags=unit -v ./test/unit/...

.PHONY: test-integ
test-integ:
	go test -tags=integration -v ./test/integration/...

.PHONY: test-race
test-race:
	go test -v -race ./...

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: fmt
fmt:
	find . -name '*.go' -not -path './internal/models/*' -not -path '*_gen.go' -print0 | xargs -0 goimports -w

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: mod
mod:
	go mod tidy -v

.PHONY: report
report: ## usage: make report type=heap|profile|block|mutex|trace
	curl -s http://127.0.0.1:9090/debug/pprof/$(type) > ./$(type).out
ifeq ($(type),trace)
	go tool trace -http=:9091 ./$(type).out
else
	go tool pprof -http=:9091 ./$(type).out
endif

.PHONY: gen
gen:
	go generate ./...

.PHONY: genORM
genORM:
	sqlboiler --wipe psql

.PHONY: migrate-status
migrate-status:
	goose -dir=./migrations postgres "postgres://postgres:12345@127.0.0.1:5432/gamesparks_db" status

.PHONY: migrate-create
create-migrate:
	goose -dir=./migrations postgres "postgres://postgres:12345@127.0.0.1:5432/gamesparks_db" create $* sql

.PHONY: migrate
migrate:
	goose -dir=./migrations postgres "postgres://postgres:12345@127.0.0.1:5432/gamesparks_db" up

.PHONY: unmigrate
unmigrate:
	goose -dir=./migrations postgres "postgres://postgres:12345@127.0.0.1:5432/gamesparks_db" down
