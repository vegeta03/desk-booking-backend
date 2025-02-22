# Commands used

1. `docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Batman -d postgres:17.3-bookworm`
2. `docker exec -it postgres17 psql -U root`
3. `docker logs postgres17`
