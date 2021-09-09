include .env

export $(shell sed 's/=.*//' .env)

$(info environment: $(GO_ENV))

.PHONY: postgres adminer migrate

build:
	@echo "🖥️  Building development container";
	docker-compose up --build --no-start

exp:
	export POSTGRESQL_URL="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432</$(POSTGRES_DB)?sslmode=disable"


run:
	@echo "🖥️  Starting the services...";
	docker-compose run --service-ports -d server

up: 
	@echo "🖥️  Starting the services...";
	docker-compose up --build



migrate:
	@echo "🗄️  Migrating database"
	migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable -path db/migrations up

remove:
	@echo "⚠️  Removing docker environment"
	docker rm postgresdb
	docker rm ginserver

restart:
	@echo "⚠️  Removing docker environment"
	docker rm postgresdb
	docker rm ginserver
	
	@echo "🖥️  Starting the services...";
	docker-compose up --build