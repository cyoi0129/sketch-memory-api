# sketch-memory-api
Sketch API using Docker

## Main Architecture
 - DB: Postgresql
 - Server: Go
 - API: Graphql

## Library
 - gqlen             https://gqlgen.com/
 - DataLoader        github.com/graph-gophers/dataloader
 - Postgres Driver   github.com/lib/pq

## gqlen Usage
### Init graph file
```sh
go run github.com/99designs/gqlgen init
```

### Generate graph file
```sh
go run github.com/99designs/gqlgen generate
```

## TBD
Solve N+ issue