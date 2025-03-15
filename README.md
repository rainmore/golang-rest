# GoLang REST example

## Database Config

Scripts to initial database and create tables

```bash
cat resources/sql/init-db.sql | docker exec -i rest-db psql -U example
cat resources/sql/init-tables.sql | docker exec -i rest-db psql -U example
```

Script to access the database
```bash
docker exec -it rest-db psql -U example
```

Script to remove volumn when remove the container
```bash
docker volume rm $(docker volume ls -q)
```