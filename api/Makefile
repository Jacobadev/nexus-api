.PHONY: migrate_up migrate_down test

DB_FILE = db/main.db
MIGRATION_UP = ./db/migrations/userCreateAndInsert.sql
MIGRATION_DOWN = ./db/migrations/dropUserTable.sql
CREATE_MODEL =./db/migrations/createUserTable.sql 

migrate_up:
	sqlite3 $(DB_FILE) < $(MIGRATION_UP)

migrate_down:
	sqlite3 $(DB_FILE) < $(MIGRATION_DOWN)

create_model:
	sqlite3 $(DB_FILE) < $(CREATE_MODEL)

test:
	@go test ./... | grep -v '?'
create_db:
	@sqlite3 $(DB_FILE) ".schema"
