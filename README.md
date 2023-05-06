# Stocker Quant

## Description
Project `Stcoker Quant` is aiming to build a quantative analysiss in Taiwan stock market.

## Migration

### Requirement
- `PostgreSQL 12 or above`
- `psql`
- `golang-migrate` ([GitHub Link](https://github.com/golang-migrate/migrate))
    - Tutorial for PostgreSQL:- [GitHub Link](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

### Related commands
- PostgreSQL localhost server
```
    sudo service postgresql start
    sudo service postgresql status
    sudo service postgresql stop
```
- PostgreSQL connection
```
    psql -h {host} -d {table_name} -U {username} -p {port}
```
- Create migrations:
```
migrate create -ext sql -dir web/migrate/data -seq init
```
- Perform migrate
```
// POSTGRESQL_URL is:
// postgres://{username}:{password}@{host}:{port}/{database_name}?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path web/migrate/data up
migrate -database ${POSTGRESQL_URL} -path web/migrate/data down
```
