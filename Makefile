version?=0.0.1
username?=devoptimus
build:
	docker build -t devoptimus/ged-client-fpm --target=fpm .
	docker build -t devoptimus/ged-client-node --target=node .
	docker build -t devoptimus/ged-client-nginx --target=nginx .
	docker build -t devoptimus/ged-client-grpc-server --target=server .

up: build
	docker compose up -d

down:
	docker compose down

push: build
	docker login -u $(username)
	
	docker push devoptimus/ged-client-fpm
	docker push devoptimus/ged-client-fpm:$(version)
	
	docker push devoptimus/ged-client-node
	docker push devoptimus/ged-client-node:$(version)

	docker push devoptimus/ged-client-nginx
	docker push devoptimus/ged-client-nginx:$(version)

	docker push devoptimus/ged-client-grpc-server
	docker push devoptimus/ged-client-grpc-server:$(version)

