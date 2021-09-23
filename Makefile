.PHONY: up-docker down-docker start-docker stop-docker \
	add-command gen-wire migrate gen-xo test clean

ADD_COMMAND_NAME = ""
APP      = ./bin/os

run:
	go build -o ${APP} ./main.go; ${APP} import

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

migrate-dry-run:
	docker exec -it -w /go/src/github.com/ph-piment/onion-scraper golang /bin/sh -c "cat ./schemas/postgres/* | psqldef -h postgres -U root -W root os --dry-run"

migrate:
	docker exec -it -w /go/src/github.com/ph-piment/onion-scraper golang /bin/sh -c "cat ./schemas/postgres/* | psqldef -h postgres -U root -W root os"

gen-xo:
	xo schema "pgsql://root:root@localhost:5432/os?sslmode=disable" -o ./app/infrastructure/dao --src templates

test:
	go clean -testcache
	go test -v ./...

clean:
	rm -rf $(APP)
