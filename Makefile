run:
	air ./cmd/web -b 0.0.0.0

up:
	docker compose up -d

down:
	docker compose down

createdb:
	docker exec -it go-starter-db-1 createdb --username=postgres --owner=postgres go-starter

dropdb:
	docker exec -it go-starter-db-1 dropdb --username=postgres -W go-starter

migrateup:
	migrate -path database/migrations -database "postgresql://postgres:password@localhost:5432/go-starter?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://postgres:password@localhost:5432/go-starter?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: run createdb dropdb up down migrateup migratedown sqlc test
