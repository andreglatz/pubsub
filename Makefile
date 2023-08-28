default: run

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...

run:
	go run main.go
