version?=0.0.1
username?=devoptimus
build:
	docker build -t devoptimus/client-fpm --target=fpm .
	docker build -t devoptimus/client-node --target=node .
	docker build -t devoptimus/client-nginx --target=nginx .

up: build
	docker compose up -d

down:
	docker compose down

push: build
	docker login -u $(username)
	
	docker push devoptimus/client-fpm
	docker push devoptimus/client-fpm:$(version)
	
	docker push devoptimus/client-node
	docker push devoptimus/client-node:$(version)

	docker push devoptimus/client-nginx
	docker push devoptimus/client-nginx:$(version)

