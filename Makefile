VERSION=#{VERSION}
EXCLUDE_THIRD_PARTY=--exclude-path third_party/errors --exclude-path third_party/google --exclude-path third_party/openapi --exclude-path third_party/validate

setup:
	go mod vendor
	go install github.com/cespare/reflex@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/improbable-eng/grpc-web/go/grpcwebproxy@latest

.PHONY: api
api:
	buf generate ${EXCLUDE_THIRD_PARTY} --path api/open-scanner-be/v1

build:
	go build -ldflags "-X main.Version=${VERSION}" -v -o bin/app-api cmd/app-api/*.go

start-local:
	reflex -r "\.(go|yaml)" -s -- sh -c "make build && ./bin/app-api -config=./files/config/local.yaml"
	
start-grpc-web:
	grpcwebproxy --backend_addr=localhost:8081 --run_tls_server=false --allow_all_origins

start-prod:
	./bin/app-api -config=./files/config/prod.yaml