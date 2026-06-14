.PHONY: help backend-fmt backend-test backend-vet frontend-install frontend-lint frontend-build compose-config compose-up compose-down db-migrations-list demo validate

help:
	@printf "Amanora developer commands\n\n"
	@printf "  make backend-fmt       Format Go backend\n"
	@printf "  make backend-test      Run Go tests\n"
	@printf "  make backend-vet       Run Go vet\n"
	@printf "  make frontend-install  Install frontend dependencies\n"
	@printf "  make frontend-lint     Run frontend lint\n"
	@printf "  make frontend-build    Build frontend\n"
	@printf "  make compose-config    Validate Docker Compose config\n"
	@printf "  make compose-up        Start local stack\n"
	@printf "  make compose-down      Stop local stack\n"
	@printf "  make db-migrations-list List SQL migrations\n"
	@printf "  make demo              Print demo walkthrough\n"
	@printf "  make validate          Run local validation suite\n"

backend-fmt:
	cd backend && gofmt -w $$(find . -name '*.go')

backend-test:
	cd backend && go test ./...

backend-vet:
	cd backend && go vet ./...

frontend-install:
	cd frontend && npm install

frontend-lint:
	cd frontend && npm run lint

frontend-build:
	cd frontend && npm run build

compose-config:
	docker compose -f deploy/docker-compose.yml config

compose-up:
	docker compose -f deploy/docker-compose.yml up -d --build

compose-down:
	docker compose -f deploy/docker-compose.yml down

db-migrations-list:
	ls -1 backend/migrations/*.sql

demo:
	@cat docs/demo-walkthrough.md

validate: backend-fmt backend-test backend-vet frontend-lint frontend-build compose-config
