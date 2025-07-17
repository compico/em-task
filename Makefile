MIGRATE_CREATE_CMD=migrate create -ext sql -dir migrations

up:
	docker-compose up -d
	docker-compose logs -f app

recreate:
	docker-compose down
	docker-compose build #--no-cache
	docker-compose up -d
	docker-compose logs -f app

migrate_up:
	docker-compose exec app /app/tmp/main migrate

run:
	TZ=UTC go run ./cmd http

.PHONY: docs
docs:
	swag init -g web/router/routes.go

.PHONY: migration help
migration:
	@if [ -z "$(name)" ]; then \
		echo "Нет параметра name"; \
		echo "Пример: make migration name=add_users_table"; \
	else \
		$(MIGRATE_CREATE_CMD) -seq $(name); \
	fi

wire_http:
	wire ./cmd/commands/http/wire.go
wire_migrate:
	wire ./cmd/commands/migrate/wire.go

download_toolchain:
	go install github.com/air-verse/air@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

init:
	@test -f ./configs/config.yaml || cp ./configs/config.yaml.dist ./configs/config.yaml
	@test -f docker-compose.yml || cp docker-compose.yml.dist docker-compose.yml
	make wire

wire: wire_http wire_migrate
