start-api:
	docker build --tag jsonrpc ./api
	docker run --rm -p 5000:5000 --name jsonrpc jsonrpc:latest