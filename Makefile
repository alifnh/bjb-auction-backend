run-test:
	go run ./cmd/api/

test:
	go test -coverpkg=./... -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | grep total
	rm coverage.out

mock:
	mockery --all --case underscore

build:
	go build -o ./build/main ./cmd/api/main.go

migrateforce:
	migrate -path ./internal/database/migration/ -database "postgres://postgres:Akundb123@localhost:5432/healthcare_db?sslmode=disable" -verbose force 1

migratedown:
	migrate -path ./internal/database/migration/ -database "postgres://postgres:Akundb123@localhost:5432/healthcare_db?sslmode=disable" -verbose down 1

migrateup:
	migrate -path ./internal/database/migration/ -database "postgres://postgres:Akundb123@localhost:5432/healthcare_db?sslmode=disable" -verbose up 1

.PHONY: run-rest test mock migrateforce migratedown migrateup