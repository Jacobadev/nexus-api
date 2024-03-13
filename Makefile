.PHONY: migrate_up migrate_down test

DB_FILE = ./api/db/main.db
MIGRATION_UP = ./migrations/userCreateAndInsert.sql
MIGRATION_DOWN = ./migrations/dropUserTable.sql

migrate_up:
	sqlite3 $(DB_FILE) < $(MIGRATION_UP)

migrate_down:
	sqlite3 $(DB_FILE) < $(MIGRATION_DOWN)

test:
	@go test ./... | grep -v '?'
create_db:
	@sqlite3 $(DB_FILE) ".schema"
