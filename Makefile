up:
	docker-compose up -d
	docker-compose logs -f app

recreate:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d
	docker-compose logs -f app

run:
	TZ=UTC go run ./cmd http

wire_http:
	wire ./cmd/commands/http/wire.go

init:
	@test -f ./configs/config.yaml || cp ./configs/config.yaml.dist ./configs/config.yaml
	@test -f docker-compose.yml || cp docker-compose.yml.dist docker-compose.yml
	go install github.com/air-verse/air@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/a-h/templ/cmd/templ@latest
	make wire

wire: wire_http