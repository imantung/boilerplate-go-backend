# Boilerplate Go Backend

Boilderplate project for golang backend

- Application
    - [X] [Go-Standards](https://github.com/golang-standards/project-layout) Project Layout
    - [x] [SOLID Principle](https://en.wikipedia.org/wiki/SOLID)
    - [x] Dependency Injection (Uber Dig)
    - [ ] Env Config
    - [X] Graceful Shutdown
    - [ ] Makefile
- API
    - [x] OpenAPI Standard 3.0
    - [x] [Echo Framework](https://echo.labstack.com/)
    - [ ] HealthCheck API
    - [ ] Custom error handler
    - [ ] Custom logger 
    - [ ] Access Control List (ACL)
    - [ ] OAuth2 Auth
    - [ ] Cache (Redis)
- Testing
    - [ ] Table Driven Test
    - [ ] Mocking with gomock
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
    - [x] Go-generate
    - [ ] Dotenv file
    - [ ] Controler (+Service) sources
    - [ ] Repository sources

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