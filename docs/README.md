# Go Services Scaffold

This is an scaffold for golang web-service.

## Get started



## Architecture



## Directory layout

| Path | Description|
| ----- | ----------------- |
| **cmd** |placement for the [cli.Commands][cli] |
| **config** | parsing of configuration file and app/packages initializing|
| **dbschema** | contains `bindata.go` with migrations and function for injecting they into db package|
| **dbschema/migrations** | must contain `.sql` files wtih db migrations|
| **docs** | project documentation|
| **models** | type definitions of db models and queriers for them|
| **internal** | some internal stuffs, such a small helpers, code for interaction with another services, etc|
| **vendor** | directory with the project dependencies|
| **workers** | in `main.go` initialization of the `routines.Cheif` with workers, subdirectories contains implementations of workers submodules|


##### Example of project layout

```text
.
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── cmd
│   └── main.go
├── config
│   ├── cfg.go
│   ├── links.go
│   ├── main.go
│   └── workers.go
├── config.yaml.tmpl
├── dbschema
│   ├── bindata.go
│   ├── migration.go
│   └── migrations
│       └── 001_base.sql
├── docs
│   └── README.md
├── main.go
├── models
│   └── main.go
└── workers
    ├── api
    │   ├── handler
    │   │   ├── get_items.go
    │   │   └── main.go
    │   ├── main.go
    │   └── middleware
    ├── cleaner
    │   └── main.go
    └── main.go
```

[cli]: https://github.com/urfave/cli