.PHONY: start-dev
start-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		-f deployments/docker/backend.docker-compose.yaml \
		-f deployments/docker/connector.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		up -d

args=
.PHONY: stop-dev
stop-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		-f deployments/docker/backend.docker-compose.yaml \
		-f deployments/docker/connector.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		down $(args)

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

.PHONY: build-backend
build-backend:
	go build -o bin/backend cmd/backend/server.go

.PHONY: build-connector
build-connector:
	go build -o bin/connector cmd/connector/server.go

version=
.PHONY: build-images
build-images:
	docker build -t backend-service -f deployments\backend\Dockerfile.backend .
	docker build -t connector-service -f deployments\connector\Dockerfile.connector .
ifdef version
	docker image tag backend-service:latest backend-service:$(version)
	docker image tag connector-service:latest connector-service:$(version)
endif
	docker image prune
