# Boilerplate Go Backend

Boilerplate project for golang backend. 

## Getting Started

### Prerequisite
- Programming Language: [Go](https://go.dev/) 
- Task Manager/Build-Tool: [GoTask](https://taskfile.dev/)
- Database Migration: [Golang-Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- Infrastructure: [Docker](https://www.docker.com/)

### Command Line
```bash
# First setup for mac (for other OS, please refer to respective documentation)
brew install go
brew install go-task
brew install golang-migrate

task --list-all   # Show available tasks
task run          # Run the project
task clean        # Clean the dev local environment

task gen-oapi        # Generate open-api controller from api-spec.yaml
task gen-dotenv      # Generate .env file from config struct
task gen-entity      # Generate entity/repo from database table
task gen-converter   # Generate converter from DAO to Entity and vice-versa

task create-migration NAME=create_some_table   # Create new migration file
```

## Overview
### Framework/Library 

- [Echo](https://echo.labstack.com/): High performance, minimalist Go web framework
- [Dig](https://github.com/uber-go/dig): A reflection based dependency injection toolkit for Go.
- [OApi-CodeGen](https://github.com/oapi-codegen/oapi-codegen): Generate Go client and server boilerplate from OpenAPI 3 specifications
- [Go-Auth2](https://github.com/go-oauth2/oauth2): OAuth 2.0 server library for the Go programming language
- [Go-Mock](https://github.com/uber-go/mock): a mocking framework for the Go programming language
- [Zerolog](https://github.com/rs/zerolog): Zero Allocation JSON Logger

### Checklist

- General
  - [x] [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
  - [x] [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) with [Uber Dig](https://github.com/uber-go/dig) -- Check the code at [internal/app/infra/di/di.go](internal/app/infra/di/di.go)
  - [x] Centralized config (env variable) -- Check the code at [internal/app/infra/config/config.go](internal/app/infra/config/config.go)
  - [x] Graceful Shutdown -- Check the code at [cmd/boilerplate-go-backend/main.go](cmd/boilerplate-go-backend/main.go)
  - [x] Health check -- Check the code at [internal/app/health.go](internal/app/health.go)
- API Server
  - [x] [OpenAPI Specification 3.0](https://swagger.io/resources/open-api/) -- Check the specification at [api/api-spec.yml](api/api-spec.yml)
    - [x] Enable [swagger-ui](https://github.com/swagger-api/swagger-ui) (`/swagger/ui`)
  - [x] [Echo Framework](https://echo.labstack.com/)  -- Check the router at [internal/app/router.go](internal/app/router.go)
    - [x] [Recovery from panic](https://echo.labstack.com/docs/middleware/recover)
    - [x] [Generate Request ID](https://echo.labstack.com/docs/middleware/request-id)
    - [x] [Protection against XSS attack and other common security threats](https://echo.labstack.com/docs/middleware/secure)
    - [x] Custom error handler
    - [ ] Server-Side Cache (Redis)
    - [ ] Rate Limiter (Redis)
- [RESTful API](https://en.wikipedia.org/wiki/REST)
  - [x] `POST` Create operation 
  - [x] `GET` Read operation 
    - [x] Single entity
    - [x] Entity List
    - [ ] Pagination
    - [ ] Search / Filtering
    - [ ] Sorting
  - [x] `UPDATE` Update operation
  - [x] `PATCH` Partially update operation
  - [x] `DELETE` Delete operation
    - [x] Soft-Delete
    - [x] Idempotent
- [Layered Architecture](https://herbertograca.com/2017/08/03/layered-architecture/)
   -  [x] Controller Layer -- Generated by [oapi-codegen](https://github.com/deepmap/oapi-codegen)
   -  [x] Business Logic Layer (Services) -- Check the code at [internal/app/service](internal/app/service)
      - [x] Validation Logics
      - [x] Trigger Database Transaction (`BEGIN`) -- Check [dbtxn](https://github.com/imantung/dbtxn) library
   -  [x] Data Access Layer (Repos) - Generated by [tools/entity-gen](tools/entity-gen)
      - [x] ORMHate Philosophy
      - [x] Query Builder (using [Squirrel](https://github.com/Masterminds/squirrel))
- Database
  - [x] PostgresSQL Database -- Check the code at [internal/app/infra/database/postgres.go](internal/app/infra/database/postgres.go)
    - [x] Connection pool 
  - [x] Audit Columns (`created_at`, `modified_at`, etc)
  - [x] Migration tool with [golang-migrate](https://github.com/golang-migrate/migrate)
  - [x] Data seeding to insert initial dummy data to database -- Check the tool at [tools/dbseed](tools/dbseed)
  - [ ] User AuditTrail (Transaction Logs)
- Security and Observability
  - [x] Basic Auth -- Check the code at [internal/app/infra/auth/basic.go](internal/app/infra/auth/basic.go)
  - [x] OAuth2 with [Go-OAuth2](https://github.com/go-oauth2/oauth2) 
    - [x] Use Postgres to store oauth-client -- Check the code at [internal/app/infra/auth/oauth_client_store.go](internal/app/infra/auth/oauth_client_store.go)
    - [ ] Use Redis to store oauth-token
    - [x] Handle authorize request and token access
    - [x] Validate bearer token and scope access -- Check the code at [internal/app/infra/auth/oauth_handler.go](internal/app/infra/auth/oauth_handler.go)
  - [x] Enable [expvar](https://pkg.go.dev/expvar) endpoint (`/debug/vars`) 
  - [x] Enable [pprof](https://pkg.go.dev/net/http/pprof) endpoint (`/debug/pprof`)
  - [x] Structured Log with [ZeroLog](https://github.com/rs/zerolog) -- Check the code at [/internal/app/infra/logger](/internal/app/infra/logger)
    - [x] Pretty format (not json) when debug enabled
    - [x] Escalate log level to `WARN` for slow API
    - [x] Append log field `pid` 
    - [x] Append log field `go_version`
    - [x] Append log field `latency` 
    - [x] Append log field `request_id`
    - [x] Append log field `user_id`
  - [ ] Open Tracing
- Code Generator
  - [x] Generate Server Interface and Controller Layer with [oapi-codegen](https://github.com/deepmap/oapi-codegen) -- Check the config at [tools/openapi-gen](tools/openapi-gen) 
  - [x] Generate Dotenv file from Config struct -- Check the tool at [tools/dotenv-gen](tools/dotenv-gen)
  - [x] Generate Entity Model and Repository Layer from Database schema -- Check the tool at [tools/entity-gen](tools/entity-gen)
  - [x] Generate Converter from DAO to Entity and vice-versa -- Check the tool at [tools/converter-gen](tools/converter-gen)
- Testing
  - [x] Table Driven Test -- Check the code at [internal/app/service/employee_svc_test.go](internal/app/service/employee_svc_test.go)
  - [x] Mocking object with [GoMock](https://github.com/golang/mock)
  - [ ] API Test Automation
- Others
  - [x] Build tool with [TaskFile](https://taskfile.dev/) (a better alternative from [GNU Make](https://www.gnu.org/software/make/))
  - [x] Dockerfile 
  - [x] Docker-compose

## Study Case: Employee Clocking System

The project use employee clocking system as the study case
- Client App API
  - [x] Clock-in
  - [x] Clock-out
- Backoffice Dashboard API
  - [x] Manage Employee
  - [x] Employee Clock History

## Notes 

- The project is OPINIONATED based on author knowledge and experience
- The project is PRAGMATIC, it use proven libraries/frameworks as much as possible without reinvented the wheel
- The project is MONOLITH BACKEND, you may need to customize code for microservices needs
- This is CONTINUOUS PROJECT, the author use it as his actual work and keep improve it
- The project is OPEN-FOR-DISCUSSION, feel free to ping the author for any feedback/question and/or open issue ticket

## FAQ

1. Echo VS Fiber? 

    [Fiber](https://github.com/gofiber/fiber) is arguably better than [Echo](https://echo.labstack.com/) but is not compatible with [net/http](https://pkg.go.dev/net/http) (it is based on [fasthttp](https://github.com/valyala/fasthttp)). We use [go-oauth2](https://go-oauth2.github.io/) who only support `net/http`, therefore we choose to use Echo instead. 

2. Pgx for postgres? 

    [Pgx](https://github.com/jackc/pgx) is a faster and more compatible postgres driver compared with [pq](https://github.com/lib/pq). There are 2 ways to use pgx: through [database/sql](https://pkg.go.dev/database/sql) and direct implementation (which offer more capability) but not compatible with `database/sql`. We want to keep compatibility with `database/sql` to give flexibility to use other library.

## Author

<iman.tung@gmail.com>


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details