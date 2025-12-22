build:
	bash ./protoc.sh
	go build -o ./bin/impact-server ./cmd/main.go

run: build
	./bin/impact-server $(inbound)

test:	
	go test ./...

test-verbose:	
	go test ./... -v