local:
	set -a; \
	. ./.env; \
	set +a; \
	go run cmd/userservice/main.go --config=./config/local.yaml

migrate_all_up:
	set -a; \
	. ./.env; \
	set +a; \
	migrate -path ./migrations -database "postgres://$$MIG_NAME:$$MIG_PASS@$$HOST:$$PORT/$$DBNAME?sslmode=$$SSLMODE" up

build_proto:
	rm -f ./proto/*.pb.go
	protoc --go_out=./proto/userservice --go_opt=paths=import --go-grpc_out=./proto/userservice --go-grpc_opt=paths=import ./proto/userservice/user.proto