postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine
createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root simplebank_db
dropdb:
	sudo docker exec -it postgres dropdb simplebank_db

migrate-up:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank_db?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank_db?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrate-up migrate-down