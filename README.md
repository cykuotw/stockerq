# Stocker Quant

## Description
Project `Stcoker Quant` is aiming to build a quantative analysiss in Taiwan stock market.

## Migration

### Requirement
- `MySQL 8.0 or above`
- `golang-migrate` ([GitHub Link](https://github.com/golang-migrate/migrate))

### Related commands
- MySQL localhost server
```
    sudo service mysql start
    sudo service mysql status
    sudo service mysql stop
```
- MySQL connection
```
    mysql -h {host} -P {port} -u {username} -p{password} {database}
```
- Create migrations:
```
migrate create -ext sql -dir web/migrate/data -seq init
```
- Perform migrate
```
// MYSQL_URL:
// mysql://{username}:{password}@tcp{{host}:{port})/{database_name}
migrate -database "${MYSQL_URL}" -path web/migrate/data up
migrate -database "${MYSQL_URL}" -path web/migrate/data down
```
