postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Batman -d postgres:17.3-bookworm

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root desk_booking

dropdb:
	docker exec -it postgres17 dropdb desk_booking

migrateup:
	migrate -path db/migration -database "postgresql://root:Batman@localhost:5432/desk_booking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:Batman@localhost:5432/desk_booking?sslmode=disable" -verbose down

.PHONY: postgres, createdb, dropdb, migrateup, migratedown