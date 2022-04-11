docker:
	docker-compose up

migrateUp:
	 migrate -path ./migrations -database postgres://postgres:123456@localhost:5432/postgres?sslmode=disable up

migrateDown:
	 migrate -path ./migrations -database postgres://postgres:123456@localhost:5432/postgres?sslmode=disable down

proto:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/users.proto

getEvans:
	go install github.com/ktr0731/evans@latest

evans:
	evans api/proto/users.proto -p 50082

run:
	go run cmd/main.go