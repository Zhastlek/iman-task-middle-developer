proto-parser:
	protoc -I ./proto --go_out=./connection_grpc --go_opt=paths=source_relative \
	--go-grpc_out=./connection_grpc --go-grpc_opt=paths=source_relative \
	parser.proto

proto-editor:
	protoc -I ./proto --go_out=./connection_grpc --go_opt=paths=source_relative \
	--go-grpc_out=./connection_grpc --go-grpc_opt=paths=source_relative \
	editor.proto

docker:
	docker-compose up

run:
	go run post-parser/cmd/main.go
	go run post-editor/cmd/main.go
	go run api/cmd/main.go

db-tables-up:
	go run database-schema/main.go

all: api parser editor

api: go run api/cmd/main.go

parser: go run post-parser/cmd/main.go

editor:	go run post-editor/cmd/main.go
