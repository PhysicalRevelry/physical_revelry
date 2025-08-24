.PHONY: setup dev build test clean

setup:
	@echo "Setting up project..."
	cd backend && go mod download
	cd mobile && flutter pub get
	docker-compose up -d

dev:
	@echo "Starting development servers..."
	docker-compose up -d
	cd backend && go run cmd/server/main.go &
	cd mobile && flutter run

build-backend:
	cd backend && go build -o bin/server cmd/server/main.go

build-mobile:
	cd mobile && flutter build apk

test-backend:
	cd backend && go test ./...

test-mobile:
	cd mobile && flutter test

clean:
	docker-compose down
	cd backend && go clean
	cd mobile && flutter clean

generate-gql:
	cd backend && go run github.com/99designs/gqlgen generate