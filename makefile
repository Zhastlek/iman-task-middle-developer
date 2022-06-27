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
	go run database-schema/main.go &
	go run post-parser/cmd/main.go &
	go run post-editor/cmd/main.go &
	go run api/cmd/main.go

db-tables-up:
	go run database-schema/main.go

pusk:
	docker image build -f api/dockerfile . -t imagename
	docker container run -p 9000:9000 -d --name api imagename

clean:
	docker system prune -a