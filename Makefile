VOTING_BINARY=votingApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d --force-recreate
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose and deletes binaries
up_build: build_voting deploy clean_build

deploy:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d --force-recreate
	@echo "Docker images built and started!"

clean_build:
	@echo "Cleaning build binaries"
	rm ./voting/${VOTING_BINARY}

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_voting: builds the voting binary as a linux executable
build_voting:
	@echo "Building voting binary..."
	cd ./voting && env GOOS=linux CGO_ENABLED=0 go build -o ${VOTING_BINARY} ./cmd/api
	@echo "Done!"
