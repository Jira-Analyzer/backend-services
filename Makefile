.PHONY: start-dev
start-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		up -d

args=
.PHONY: stop-dev
stop-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		down ${args}

.PHONY: unit-test
unit-test:
	go test -cover ./... -tags=unit  -covermode=count -coverprofile=coverage.out
	go tool cover -func=coverage.out -o=coverage.out
	gobadge -filename=coverage.out

.PHONY: gogen
gogen:
	go generate ./...

.PHONY: swag-backend
swag-backend:
	swag init -g cmd/backend/server.go --exclude internal/handler/connector --output docs/backend-service --parseInternal


.PHONY: swag-connector
swag-connector:
	swag init -g cmd/connector/server.go --exclude internal/handler/backend --output docs/connector-service --parseInternal
