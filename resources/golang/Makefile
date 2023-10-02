gen:
	protoc \
	--proto_path=proto \
	--go_out=pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=pb \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto

clean:
	rm -f pb/*.pb.go

run:
	go run main.go

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

docker-build:
	docker build -t devoptimus/ged-client-protobuf --target=server .

docker-up: docker-down
	docker compose up -d --build

docker-down:
	docker compose down