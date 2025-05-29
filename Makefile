DB_URL=postgresql://postgres:postgres@localhost:5433/chatapp?sslmode=disable

network:
	docker network create chatapp

postgres:
	docker run --name postgres --network chatapp -p 5433:5433 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres chatapp

dropdb:
	docker exec -it postgres dropdb -U postgres chatapp

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
	@echo "Migration file created in db/migration"
	@echo "Please add the SQL commands in the up and down files."
	@echo "Then run 'make migrateup' to apply the migration."
	@echo "And 'make migratedown' to revert the migration."
	@echo "Example: make new_migration name=user_table"

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration