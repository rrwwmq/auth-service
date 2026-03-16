include .env
export

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down

up:
	docker-compose up -d

down:
	docker-compose down

run:
	go run ./cmd/main.go

