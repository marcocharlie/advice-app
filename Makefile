test-api:
	docker build --tag jsonrpctest . -f ./api/test.dockerfile
	docker run --rm jsonrpctest:latest

start-api:
	docker build --tag jsonrpc ./api
	docker run --rm -p 5000:5000 --name jsonrpc jsonrpc:latest