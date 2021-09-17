.PHONY: up-docker down-docker start-docker stop-docker \
	add-command gen-wire migrate gen-xo test clean

ADD_COMMAND_NAME = ""
APP      = ./bin/os

up-docker:
	docker-compose up -d

down-docker:
	docker-compose down

start-docker:
	docker-compose start

stop-docker:
	docker-compose stop

# make add-command ADD_COMMAND_NAME=hoge
add-command:
	cobra add $(ADD_COMMAND_NAME)

gen-wire:
	wire ./cmd/di

migrate:
	migrate -database 'postgres://root:root@postgres:5432/os?sslmode=disable' -path ./migrations/ up

gen-xo:
	xo schema "pgsql://root:root@localhost:5432/os?sslmode=disable" -o ./app/infrastructure/dao --src templates

test:
	go clean -testcache
	go test -v ./...

clean:
	rm -rf $(APP)
