test-api:
	docker-compose -f ./api/docker-compose.test.yml up --build --abort-on-container-exit

start-api:
	docker-compose -f ./api/docker-compose.yml up --build --force-recreate -d

stop-api:
	docker-compose -f ./api/docker-compose.yml stop

start-client:
	docker-compose -f ./client/docker-compose.yml up --build --force-recreate -d

stop-client:
	docker-compose -f ./client/docker-compose.yml stop