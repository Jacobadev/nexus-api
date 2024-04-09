.PHONY: migrate_up migrate_down test run

DB_FILE = db/main.db
MIGRATION_UP = ./db/migrations/userCreateAndInsert.sql
MIGRATION_DOWN = ./db/migrations/dropUserTable.sql
CREATE_MODEL =./db/migrations/createUserTable.sql 

run: 
	@go run cmd/main.go
create:
	psql -Upostgres -dnexus -af migrations/usernew.sql
seed:
	psql -Upostgres -dnexus -af migrations/seed.sql
drop:
	psql -Upostgres -dnexus -af migrations/drop.sql
test:
	@go test
