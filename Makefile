server:
	go run cmd/app/main.go

grpc-redis:
	echo -e "STORAGE_TYPE=redis\nSERVER_TYPE=grpc" > configs/types.env
	docker compose -f docker-compose.redis.yaml up --build

grpc-postgres:
	echo -e "STORAGE_TYPE=postgres\nSERVER_TYPE=grpc" > configs/types.env
	docker compose -f docker-compose.psql.yaml up --build

http-redis:
	echo -e "STORAGE_TYPE=redis\nSERVER_TYPE=http" > configs/types.env
	docker compose -f docker-compose.redis.yaml up --build

http-postgres:
	echo -e "STORAGE_TYPE=postgres\nSERVER_TYPE=http" > configs/types.env
	docker compose -f docker-compose.psql.yaml up --build

tests:
	go test -v ./...