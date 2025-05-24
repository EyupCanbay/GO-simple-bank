postgres:
	docker run --name postgres17-alpine -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine

createdb:
	docker exec -it postgres17-alpine createdb --username postgres --owner postgres simple_bank

dropdb:
	docker exec -it postgres17-alpine dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...  

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migratedown migrateup sqlc test server mock