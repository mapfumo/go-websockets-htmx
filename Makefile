run:
	go run cmd/main.go

test:
	go test -v ./...

coverage:
	go test -cover -v ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...