# Tnderlike
Tinder-like backend REST API

## How to Run
- Setup environment, Create environment variables file by copying the example in cmd/tnderlike/config/server/ or generate by using the envgen script
    ```make envgen```

- Run by using Docker or see [Run](#run) section below
    ```docker-compose up```

## Structure
```
.
|-cmd
|---tnderlike (location of entrypoint/function main)
|-----config (config for the service)
|-------server
|-docker (docker files)
|---db
|-----migration
|-internal (service functionality)
|---api (service/router initialization)
|-----server
|-------router (router setup and initialization)
|---config (config handler)
|---database (database handler)
|---domain (service functionality)
|-----constant (constant variable)
|-----controller
|-------auth
|-------ping
|-----model
|-------auth
|-------common
|-----repository
|-------auth
|-----service
|-------auth
|-------ping
|---lib (lib used for internal)
|-----account
|-----crypto
|-------aes
|-------argon
|-----paseto
|---middleware
|-----auth
|-logs
|-mocks (mocks used for unit testing)
|---internal_
|-----domain
|-------repository
|-------service
|-----lib
|-------paseto
|-pkg (common libs)
|---json
|---logger
|---response
|-scripts (scripts)
```

## Run

### Run using Docker

Prerequisites:
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

Command:
```docker-compose up```

### Run without Docker
Prerequisites:
- Setup Postgresql
- Golang

Command:
```make run```

### Expected Output on Running

```
INF POST : /login
INF GET : /ping
INF POST : /register
INF Running at :80
```


## Test
### Run unit tests
Run unit tests by running ```go test ./...```
```make test```

### Run API tests
Test API by executing scripts/testapis.sh
```make testapi```
