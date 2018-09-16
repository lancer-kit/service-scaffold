# DB 

DB  is a package with helpers and wrappers for interacting with PostgreSQL using the [`squirrel.Builder`](https://github.com/Masterminds/squirrel).

## Package types

- `Transactional` is the interface for representing a db connector/query builder that support database transactions.
- `SQLConn` is a connector for execution of queries to the PostgreSQL database.
- `PageQuery` is the structure for building query with pagination
- `Table` is the basis for implementing Querier for some model or table.

# SQLConn

This is a database connector. **SQLConn** supports *database transactions*.

- `Clone` returns sanitized copy of SQLConn instance backed by the same context and db. The result will not be bound to any transaction that the source is currently within.
- `Get` casts given `squirrel.Sqlizer` to raw SQL string with arguments and execute `GetRaw`.
- `Exec` casts given `squirrel.Sqlizer`  to raw SQL string with arguments and execute `ExecRaw`
- `Select` casts given `squirrel.SelectBuilder` to raw SQL string with arguments and execute `SelectRaw`.

- `GetRaw` executes given SQL Select and returns only **one** record.
- `ExecRaw` executes any given SQL query.
- `SelectRaw` executes given SQL Select and returns **all** records.

# PageQuery

# Table

# Usage 

- Connect and execute some query:

```go
package main

import (
    "log"
    "gitlab.inn4science.com/gophers/service-kit/db"
)

func main() {
    // initialize SQLConn singleton
    err := db.Init("postgres://user:user@localhost/user?sslmode=disable", nil)
    if err != nil {
        panic(err)
    }
    sqlConn := db.GetConnector()
    err = sqlConn.ExecRaw(`CREATE TABLE IF NOT EXIST users(
    id SERIAL, name VARCHAR(64), email VARCHAR(64), age INTEGER)`, nil)
    if err != nil {
        panic(err)
    }
}
```

- [Full Querier implementation](../internal/examples/db.go)