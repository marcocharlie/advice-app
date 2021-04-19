test-backend:
	docker-compose -f ./api/docker-compose.test.yml up --build --abort-on-container-exit

start-app:
	docker-compose -f ./docker-compose.yml up --build --force-recreate -d

stop-app:
	docker-compose -f ./docker-compose.yml stop