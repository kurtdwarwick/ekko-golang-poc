build:
	bash ./protoc.sh
	go build -o ./bin/organisation-server ./cmd/main.go

run: build
	export OTEL_EXPORTER_OTLP_INSECURE="true"
	./bin/organisation-server

mockgen:
	mockgen -source=internal/features/onboard/core/data/access/organisation_dao.go -destination=internal/features/onboard/core/data/access/organisation_dao_mock.go -package=access

test: mockgen	
	go test ./...

test-verbose: mockgen	
	go test ./... -v