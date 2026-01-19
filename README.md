## Financial Control Engine

### What is?
The Financial Control Engine is an internal Go service that implements the business logic and data layer for financial control operations. It centralizes domain rules, persistence (DTOs / repositories), and migrations so other internal components can interact with a single, consistent source of truth for financial data.

Key points
- Purpose: encapsulate business logic, data models and persistence for all financial control concerns.
- Audience: internal services and authenticated clients inside the platform (not public-facing).
- Responsibilities: domain validation, transactional operations, DB access, DTOs, schema migrations and audit logging.
- Not responsible for authentication/authorization: the service expects requests to arrive already authenticated and to include the user identity (e.g., user id in a validated token). It will *
read
* the user id from the request context/token but does not validate credentials itself.

### Instalation & development build

#### 1 - create a .env file like the example.env

```env
FINANCIAL_CONTROL_DATABASE_PORT=5432
FINANCIAL_CONTROL_DATABASE_NAME=financialcontrol
FINANCIAL_CONTROL_DATABASE_USER=docker
FINANCIAL_CONTROL_DATABASE_PASSWORD=docker
FINANCIAL_CONTROL_DATABASE_HOST=localhost
FINANCIAL_CONTROL_CSRF_KEY=NhkEXjyS5ms3k7vNQ5fbk2Ffv0OIuQs6
```

#### 2 - run the docker compose image to create the database

```shell
docker compose up -d
```

#### 3 - install the dependencies

```shell
go mod download
go mod tidy
```

#### 4 - run the /terndotenv main.go to run the migrations using [tern](https://github.com/jackc/tern)

```shell
go run ./cmd/terndotenv
```

#### 4.1 - run the command below to rollback migrations in linux
```shell
 export $(cat example.env | xargs) && tern migrate --migrations ./internal/store/pgstore/migrations --config ./internal/store/pgstore/migrations/tern.conf --destination=0
```

#### 5 - run the command below for dev mode using [air](https://github.com/air-verse/air)

```shell
air
```