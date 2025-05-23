postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres17 dropdb --username=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down		
sqlc:
	docker run --rm -v "$(CURDIR):/src" -w /src sqlc/sqlc generate
test:
	go test -v -cover ./...	
mock:
	mockgen -destination db/mock/store.go -package mockdb  github/heimaolst/simplebank/db/sqlc Store	
server:
	go run main.go


.PHONY: createdb dropdb	postgres migrateup migratedown sqlc test server mock