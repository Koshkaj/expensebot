install:
	@go mod download

run:
	@go run main.go

env:
	@cp .env.example .env

mongo:
	@docker run -d --name expense_mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=myuser -e MONGO_INITDB_ROOT_PASSWORD=mypass mongo

test:
	@ go test -v ./...

.PHONY: run mongo env test install