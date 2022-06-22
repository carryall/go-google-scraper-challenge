![Staging test](https://github.com/carryall/go-google-scraper-challenge/actions/workflows/test.yml/badge.svg?branch=develop)![Staging deployment](https://github.com/carryall/go-google-scraper-challenge/actions/workflows/deploy.yml/badge.svg?branch=develop)![Staging](https://pyheroku-badge.herokuapp.com/?app=google-scraper-staging&style=flat)

![Production test](https://github.com/carryall/go-google-scraper-challenge/actions/workflows/test.yml/badge.svg?branch=main)![Production deployment](https://github.com/carryall/go-google-scraper-challenge/actions/workflows/deploy.yml/badge.svg?branch=main)![Production](https://pyheroku-badge.herokuapp.com/?app=google-scraper-web&style=flat)

## Introduction

A project for Nimble Go Internal Certification on Web

[Staging](https://google-scraper-staging.herokuapp.com)
[Production](https://google-scraper-web.herokuapp.com)

## Project Setup

### Prerequisites

- [Go - 1.18](https://golang.org/doc/go1.18) or newer

- [Node - 16](https://nodejs.org/en/)

### Development

#### Create an ENV file

To start the development server, `.env` file must be created.

- Copy `.env.example` file and rename to `.env`

#### Build dependencies

- [`air`](https://github.com/cosmtrek/air) is used for live reloading

- [`goose`](https://github.com/pressly/goose) is used for database migration.

- [`forego`](https://github.com/ddollar/forego) manages Procfile-based applications.

They need to be built as a binary file in `$GOPATH`.

```make
make install-dependencies
```

#### Start development server

```make
make dev
```

The application runs locally at http://localhost:8080

### Test

Execute all unit tests:

```make
make test
```

### Migration

#### Create migration

```make
make migration/create MIGRATION_NAME={migration name}
```

#### List the migration status

```make
make migration/status
```

#### Migrate the database

```make
make db/migrate
```

#### Rollback the migration

```make
make db/rollback
```
