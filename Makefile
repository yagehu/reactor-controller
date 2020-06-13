POSTGRES_DB = reactor-controller
POSTGRES_USER = reactor
POSTGRES_PASSWORD = reactor
POSTGRES_PORT = 5432
POSTGRES_URL = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

.PHONY: db-up
db-up:
	@echo "Bringing up Postgres Docker container..."
	@POSTGRES_DB=$(POSTGRES_DB) \
		POSTGRES_USER=$(POSTGRES_USER) \
		POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		docker-compose up -d \
		> /dev/null 2>&1
	@echo "$$(tput setaf 2)Done.$$(tput sgr0)"
	@echo "Waiting for database to be available..."
	@while ! pg_isready \
		--dbname=$(POSTGRES_DB) \
		--host=localhost \
		--port=$(POSTGRES_PORT) \
		--username=$(POSTGRES_USER) \
		> /dev/null 2>&1 \
	; do \
		sleep 1; \
	done
	@echo "$$(tput setaf 2)Done.$$(tput sgr0)"
	@echo "Performing migration..."
	@migrate \
		-database $(POSTGRES_URL) \
		-path db/migrations \
		up
	@echo "$$(tput setaf 2)Done.$$(tput sgr0)"


.PHONY: db-down
db-down:
	@echo "Tearing down Postgres Docker container..."
	@docker-compose down > /dev/null 2>&1
	@echo "$$(tput setaf 2)Done.$$(tput sgr0)"

.PHONY: migrate-down
migrate-down:
	@migrate \
		-database $(POSTGRES_URL) \
		-path db/migrations \
		down