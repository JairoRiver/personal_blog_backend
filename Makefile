ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root my_blog

dropdb:
	docker exec -it postgres13 dropdb my_blog

migrateup:
	docker run --rm -v "$(ROOT_DIR)/pkg/db/migration":/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/my_blog?sslmode=disable" -verbose up

migratedown:
	docker run --rm -v "$(ROOT_DIR)/pkg/db/migration":/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/my_blog?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(ROOT_DIR):/src -w /src kjconroy/sqlc generate

server:
	go run cmd/api/main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server