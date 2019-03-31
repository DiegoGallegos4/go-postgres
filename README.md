## GO Starter template

This is a starter project template that contains the following:

- [JSON schema validation](https://github.com/go-validator/validator)
- [Logging](https://github.com/uber-go/zap)
- [Postgres ORM](https://github.com/jinzhu/gorm)
- [API Routing](https://github.com/go-chi/chi)
- Modules (go mod)

### Getting Started on your own environment (Go>1.11)

Build project:

`GO111MODULE=on go build`

Run project:

`./api`

### Getting Started with Docker

Build images:

`make build-dev`

Run images:

`make up`
