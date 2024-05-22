# Boilerplate Go Backend

Boilerplate project for golang backend

- Application
    - [X] [Go-Standards](https://github.com/golang-standards/project-layout) Project Layout
    - [x] [SOLID Principle](https://en.wikipedia.org/wiki/SOLID)
    - [x] Dependency Injection ([Uber Dig](https://github.com/uber-go/dig)) -- Check the code at [internal/app/infra/di/di.go](internal/app/infra/di/di.go)
    - [x] Centralized config (env variable) -- Check the code at [internal/app/infra/config.go](internal/app/infra/config.go)
    - [x] Graceful Shutdown -- Check the code at [cmd/boilerplate-go-backend/main.go](cmd/boilerplate-go-backend/main.go)
    - [ ] Makefile
- API
    - [x] OpenAPI Standard 3.0 -- check the API Specification at [api/api-spec.yml](api/api-spec.yml)
    - [x] [Echo Framework](https://echo.labstack.com/)
    - [ ] HealthCheck API
    - [ ] Custom error handler
    - [ ] Custom logger handler
    - [x] OAuth2 with [Go-OAuth2](https://github.com/go-oauth2/oauth2) -- Check code at [internal/app/infra/oauth.go](internal/app/infra/oauth.go)
    - [ ] Access Control List (ACL) with [Casbin](https://github.com/casbin/casbin)
    - [ ] Server-Side Cache (Redis)
- Testing
    - [ ] Table Driven Test
    - [ ] Test Automation
- Database
    - [x] PostgresSQL
    - [x] Connection pool -- Check the code at [internal/app/infra/postgres.go](internal/app/infra/postgres.go)
    - [ ] ORMHate Philosophy
    - [ ] Query Builder (using Squirrel)
    - [ ] Database Transaction in Service Layer
    - [ ] Database migration script (gomigrate)
- Docker
    - [ ] Dockerfile 
    - [ ] Dockercompose to run system component (database, redis)
- Code Generator
    - [x] Open API Stub Server using [oapi-codegen](github.com/deepmap/oapi-codegen) -- Check the go-generate at [internal/app/infra/server.go](internal/app/infra/server.go) 
    - [ ] Mocking using mockgen
    - [x] Dotenv file -- Check the tool at [tools/dotenv_gen](tools/dotenv_gen/README.md)
    - [ ] Controler (+Service) code snippet
    - [ ] Repository code snippet

## Use Case

The project use employee clocking system as use case
- For Mobile App 
    - Clock-in
    - Clock-out
- For Backoffice Dashboard
    - Employee Clock History
    - Manage User

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details