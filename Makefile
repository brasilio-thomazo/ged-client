build:
	docker build -t devoptimus/ged-client-cli --target=cli .
	docker build -t devoptimus/ged-client-fpm --target=fpm .
	docker build -t devoptimus/ged-client-nginx --target=nginx .
	docker build -t devoptimus/ged-grpc-server --target=grpc .
