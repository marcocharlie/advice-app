test-backend:
	docker-compose -f ./api/docker-compose.test.yml up --build --abort-on-container-exit

start-backend:
	docker-compose -f ./api/docker-compose.yml up --build --force-recreate -d

stop-backend:
	docker-compose -f ./api/docker-compose.yml stop

start-frontend:
	docker-compose -f ./app/docker-compose.yml up --build --force-recreate -d

stop-frontend:
	docker-compose -f ./app/docker-compose.yml stop