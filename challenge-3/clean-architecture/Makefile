.PHONY: dependency_generation migrate_up migrate_down grpc_generate_orders graphql_update_schema

migrate_up:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrate_down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

dependency_generation:
	cd ./service_container && go run github.com/google/wire/cmd/wire

grpc_generate_orders:
	export PATH="$$PATH:$(go env GOPATH)/bin" & protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

graphql_update_schema:
	go run github.com/99designs/gqlgen generate