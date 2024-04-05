.PHONY: migrate_up migrate_down test run

DB_FILE = db/main.db
MIGRATION_UP = ./db/migrations/userCreateAndInsert.sql
MIGRATION_DOWN = ./db/migrations/dropUserTable.sql
CREATE_MODEL =./db/migrations/createUserTable.sql 

run: 
	@go run cmd/main.go
migrate_up:
	sqlite3 $(DB_FILE) < $(MIGRATION_UP)

migrate_down:
	sqlite3 $(DB_FILE) < $(MIGRATION_DOWN)

create_model:
	sqlite3 $(DB_FILE) < $(CREATE_MODEL)

test:
	@go test ./test -count=1
create_db:
	@sqlite3 $(DB_FILE) ".schema"
