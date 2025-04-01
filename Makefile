include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -database $(DB_ADDR) -path $(MIGRATIONS_PATH) up

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_ADDR) -path $(MIGRATIONS_PATH) down

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) force $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

# so just run make migrate-force {version}, and it would clean the database
# migrate create -seq -ext sql -dir .\cmd\migrate\migrations\ alter_comments_cascade
# migrate -path ./cmd/migrate/migrations -database postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable up