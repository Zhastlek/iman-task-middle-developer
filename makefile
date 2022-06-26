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
	go run post-editer/cmd/main.go
	go run api/cmd/main.go

db-tables-up:
	database-schema/main.go
