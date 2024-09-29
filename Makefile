build:
	@go build -o bin/uninalerts_goapi

run: build
	@./bin/uninalerts_goapi

test:
	@go test -v ./...