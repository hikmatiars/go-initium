run-http:
	go run cmd/http/main.go

swag:
	swag init -g cmd/api/main.go

test:
	CGO_ENABLED=1 go test -race -short -count=1 ./... -gcflags=all=-l

test-coverage:
	CGO_ENABLED=1 go test -cover -race -short -count=1 -coverprofile=coverage.out ./... -gcflags=all=-l
	go tool cover -html=coverage.out