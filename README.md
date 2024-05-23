# Boilerplate Go Backend

Boilerplate project for golang backend

- Application
    - [x] [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
    - [x] [SOLID Principle](https://en.wikipedia.org/wiki/SOLID)
    - [x] [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) with [Uber Dig](https://github.com/uber-go/dig) -- Check the implementation at [internal/app/infra/di/di.go](internal/app/infra/di/di.go)
    - [x] Centralized config (env variable) -- Check the implementation at [internal/app/infra/config.go](internal/app/infra/config.go)
    - [x] Graceful Shutdown -- Check the implementation at [cmd/boilerplate-go-backend/main.go](cmd/boilerplate-go-backend/main.go)
- API
    - [x] [OpenAPI Specification 3.0](https://swagger.io/resources/open-api/) -- check the specification at [api/api-spec.yml](api/api-spec.yml)
    - [x] [Echo Framework](https://echo.labstack.com/)
    - [ ] HealthCheck API
    - [ ] Custom error handler
    - [ ] Custom logger handler
    - [x] OAuth2 with [Go-OAuth2](https://github.com/go-oauth2/oauth2) -- Check the implementation at [internal/app/infra/oauth.go](internal/app/infra/oauth.go)
    - [ ] Access Control List (ACL) with [Casbin](https://github.com/casbin/casbin)
    - [ ] Server-Side Cache (Redis)
    - [ ] Audit Trail
- Testing
    - [ ] Table Driven Test
    - [ ] Test Automation
- Database
    - [x] PostgresSQL Database
    - [x] Connection pool -- Check the implementation at [internal/app/infra/postgres.go#27](internal/app/infra/postgres.go#27)
    - [ ] Data Access Layer (DAL) / Repository Pattern
    - [ ] ORMHate Philosophy
    - [ ] Query Builder (using Squirrel)
    - [ ] Transaction in Service Layer
- Code Generator
    - [x] Open API Stub Server using [oapi-codegen](github.com/deepmap/oapi-codegen) -- Check the go-generate at [internal/app/infra/server.go](internal/app/infra/server.go) 
    - [ ] Mocking using mockgen
    - [x] Dotenv file -- Check the tool at [tools/dotenv_gen](tools/dotenv_gen/README.md)
    - [ ] Controler (+Service) template
    - [ ] Repository template
- Others
    - [ ] Makefile
    - [ ] Database migration with gomigrate
    - [ ] Dockerfile 
    - [ ] Dockercompose to run system component (database, redis)

## Use Case

The project use employee clocking system as use case
- For Mobile App 
    - Clock-in
    - Clock-out
- For Backoffice Dashboard
    - Employee Clock History
    - Manage User

## Notes 

- The project is OPINIONATED based on author knowledge and experience
- The project is PRAGMATIC, it use the other popular/proven go library as much as possible without reinvented the wheel
- The project is MONOLITH BACKEND, you may need to customize it for microservices 
- This is CONTINUOUS PROJECT, the author use the project as his actual project reference and keep improve it
- The project is OPEN-FOR-DISCUSSION, feel free to ping the author for any feedback/question and/or open issue ticket

## FAQ

1. Echo VS Fiber? [Fiber](https://github.com/gofiber/fiber) is faster and more popular web framework compared than Echo (although not signification). The caveat is Fiber based on fasthttp and not compatible with net/http which cucumbersome when use other project. In this project case, we are using [go-oauth2](https://github.com/go-oauth2/oauth2). Check the benchmark at [here](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-how-much-performance-difference-is-really-there-for-a-real-world-use-1ed29d6a3e4d).
2. Pgx for postgres? [Pgx](https://github.com/jackc/pgx) is faster and more compatible postgres driver compare with [pq](https://github.com/lib/pq). There are 2 way to use pgx: through database/sql and direct implementation (which offer more capability). The problem with 2nd approach (direct implementation) is the interface is not compatible with database/sql which is issue when want to use other library (in our case, it is query builder).

## Author

<iman.tung@gmail.com>


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details