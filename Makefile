init:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen init

generate:
# go run github.com/99designs/gqlgen
	go get github.com/99designs/gqlgen/cmd
# && go run ./internal/models/mutateHook/main.go
	go run github.com/99designs/gqlgen

run:
	docker-compose up -d

get:
	go get .

test:
	go run .

log:
	docker logs bankless-gql -f

up:
	docker-compose up -d --remove-orphans

down:
	docker compose down --remove-orphans

clean: down
	docker system prune -a --volumes