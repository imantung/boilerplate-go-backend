# Boilerplate Go Backend

Boilerplate project for golang backend. 

- General
  - [x] [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
  - [x] [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) with [Uber Dig](https://github.com/uber-go/dig) -- Check the code at [internal/app/infra/di/di.go](internal/app/infra/di/di.go)
  - [x] Centralized config (env variable) -- Check the code at [internal/app/infra/config/config.go](internal/app/infra/config/config.go)
  - [x] Graceful Shutdown -- Check the code at [cmd/boilerplate-go-backend/main.go](cmd/boilerplate-go-backend/main.go)
- API Server
  - [x] [OpenAPI Specification 3.0](https://swagger.io/resources/open-api/) -- Check the specification at [api/api-spec.yml](api/api-spec.yml)
    - [x] Enable [swagger-ui](https://github.com/swagger-api/swagger-ui) (`/swagger/ui`)
  - [x] [Echo Framework](https://echo.labstack.com/)  -- Check the code at [internal/app/app.go](internal/app/app.go)
    - [x] [Recovery from panic](https://echo.labstack.com/docs/middleware/recover)
    - [x] [Generate Request ID](https://echo.labstack.com/docs/middleware/request-id)
    - [x] [Protection against XSS attack and other common security threats](https://echo.labstack.com/docs/middleware/secure)
    - [ ] Custom error handler
    - [ ] Server-Side Cache (Redis)
    - [ ] Rate Limiter (Redis)
  - [ ] Resful API
    - [ ] Create operation
    - [ ] Read operation
    - [ ] Update operation
    - [ ] Delete operation
      - [ ] Soft-Delete
      - [ ] Idempotent
- Security and Observability
  - [x] Basic Auth -- Check the code at [internal/app/infra/auth/basic.go](internal/app/infra/auth/basic.go)
  - [x] OAuth2 with [Go-OAuth2](https://github.com/go-oauth2/oauth2) -- Check the code at [internal/app/infra/auth/oauth.go](internal/app/infra/auth/oauth.go)
    - [x] Handle authorize request
    - [x] Handle token request
    - [x] Validate bearer token
    - [ ] Validate scope access
  - [x] Health check -- Check the code at [internal/app/app.go#73](internal/app/app.go#73)
  - [x] Enable [expvar](https://pkg.go.dev/expvar) endpoint (`/debug/vars`) 
  - [x] Enable [pprof](https://pkg.go.dev/net/http/pprof) endpoint (`/debug/pprof`)
  - [x] Structured Log with [ZeroLog](https://github.com/rs/zerolog) -- Check the code at [/internal/app/infra/logger/logger.go](/internal/app/infra/logger/logger.go)
    - [x] Implemented to Echo with [Lecho](https://github.com/ziflex/lecho) 
    - [x] Pretty format (not json) when debug enabled
    - [x] Append log field `pid` 
    - [x] Append log field `go_version`
    - [x] Append log field `request_id`
    - [ ] Append log field `user_id`
    - [x] Escalate log level for slow request
  - [ ] Tracing
- Database
  - [x] PostgresSQL Database -- Check the code at [internal/app/infra/database/postgres.go](internal/app/infra/database/postgres.go)
    - [x] connection pool 
  - [ ] Data Access Layer (DAL) / Repository Pattern
    - [ ] ORMHate Philosophy
    - [ ] Query Builder (using Squirrel)
    - [ ] Database Transaction (`BEGIN`) from Business Logic Layer
  - [ ] Audit Columns (`created_at`, `modified_at`, etc)
- - [ ] User Audit Trail / Transaction Logs
- Code Generator
  - [x] Open API Stub Server using [oapi-codegen](github.com/deepmap/oapi-codegen) -- Check the go-generate at [internal/app/app.go](internal/app/app.go) 
  - [ ] Object Mocking using [gomock](https://github.com/uber-go/mock)
  - [x] Generate Dotenv file -- Check the tool at [tools/dotenv_gen](tools/dotenv_gen/README.md)
  - [ ] Generate Controler (+ Service) 
  - [ ] Generate Repository (+ SQL) 
  - [ ] Generate Entity Model from Database schema
- Testing
  - [ ] Table Driven Test
  - [ ] Test Automation
- Others
  - [x] Build tool with [TaskFile](https://taskfile.dev/) (a better alternative from [GNU Make](https://www.gnu.org/software/make/))
  - [x] Dockerfile 
  - [x] Docker-compose
  - [ ] Database migration with [go-migrate](https://github.com/golang-migrate/migrate)

## Study Case: Employee Clocking System

The project use employee clocking system as the study case
- For Client App 
  - [ ] Clock-in
  - [ ] Clock-out
- For Backoffice Dashboard
  - [ ] Manage Employee
  - [ ] Employee Clock History

## Notes 

- The project is OPINIONATED based on author knowledge and experience
- The project is PRAGMATIC, it use the other popular/proven go library as much as possible without reinvented the wheel
- The project is MONOLITH BACKEND, you may need to customize code for microservices needs
- This is CONTINUOUS PROJECT, the author use this as his actual project reference and keep improve it
- The project is OPEN-FOR-DISCUSSION, feel free to ping the author for any feedback/question and/or open issue ticket

## FAQ

1. Echo VS Fiber? 

    [Fiber](https://github.com/gofiber/fiber) is a popular and (slightly) faster web framework compared to [Echo](https://echo.labstack.com/). The caveat is that fiber is based on [fasthttp](https://github.com/valyala/fasthttp) and not compatible with [net/http](https://pkg.go.dev/net/http) which is cumbersome if we use other net/http based project. In our case, we are using [go-oauth2](https://github.com/go-oauth2/oauth2) to generate net/http server interface. Check the [benchmark between Fiber and Echo](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-how-much-performance-difference-is-really-there-for-a-real-world-use-1ed29d6a3e4d).

2. Pgx for postgres? 

    [Pgx](https://github.com/jackc/pgx) is a faster and more compatible postgres driver compared with [pq](https://github.com/lib/pq). There are 2 ways to use pgx: through [database/sql](https://pkg.go.dev/database/sql) and direct implementation (which offer more capability) but not compatible with `database/sql`. We want to keep compatibility with `database/sql` to give flexibility to use other library.

## Author

<iman.tung@gmail.com>


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details