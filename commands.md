# Commands used

1. `docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Batman -d postgres:17.3-bookworm`
2. `docker exec -it postgres17 psql -U root`
3. `docker logs postgres17`
4. `migrate create -ext sql -dir db/migration -seq init_schema`
5. `docker exec -it postgres17 /bin/sh`
6. `createdb --username=root --owner=root desk_booking`
7. `psql -U root desk_booking`
8. `dropdb desk_booking`
9. `docker exec -it postgres17 createdb --username=root --owner=root desk_booking`
10. `docker exec -it postgres17 psql -U root desk_booking`
