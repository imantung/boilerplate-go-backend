# Boilerplate Go Backend

Boilderplate project for golang backend

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
    - [ ] Custom logger 
    - [ ] Access Control List (ACL)
    - [ ] OAuth2 Auth
    - [ ] Server-Side Cache (Redis)
- Testing
    - [ ] Table Driven Test
    - [ ] Test Automation
- Database
    - [ ] Database migration script (gomigrate)
    - [ ] ORMHate
    - [x] Connection pool
    - [ ] Database Transaction in Service Layer
    - [ ] Query Builder (using Squirrel)
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
    - Login/Logout
    - Clock-in
    - Clock-out
- For Backoffice Dashboard
    - Login/logout
    - Employee Clock History
    - Manage User

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details