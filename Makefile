.PHONY: run test fmt migrate-up migrate-down


run:
go run ./cmd/api


test:
go test ./...


fmt:
gofmt -s -w .


migrate-up:
migrate -path ./migrations -database $$DATABASE_URL up


migrate-down:
migrate -path ./migrations -database $$DATABASE_URL down 1