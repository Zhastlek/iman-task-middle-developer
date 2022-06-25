proto-parser:
	api/ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/parser.proto

docker:
    docker-compose up

run:
    go run post-parser/cmd/main.go
    go run post-editer/cmd/main.go
    go run api/cmd/main.go

db-tables-up:
    go run database-schema/main.go
