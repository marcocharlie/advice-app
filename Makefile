test-api:
	docker build --tag jsonrpctest . -f ./api/test.dockerfile
	docker run --rm jsonrpctest:latest

start-api:
	docker-compose -f ./api/docker-compose.yml up --build --force-recreate -d

stop-api:
	docker-compose -f ./api/docker-compose.yml stop