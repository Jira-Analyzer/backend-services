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
