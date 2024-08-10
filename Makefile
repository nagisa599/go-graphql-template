genType:
	gqlgen generate --config ./graphql/gqlgen.yml

bl-local:
	cd ./build && docker compose build --no-cache

up-local:
	cd ./build && docker compose up

backend:
	docker exec -it grahql-backend bash

db: 
	docker exec -it mysql-db bash