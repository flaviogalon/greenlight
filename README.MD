# Greenlight
JSON API for retrieving and managing information about movies.

This is an educational project developed following the book [Let's Go Further](https://lets-go-further.alexedwards.net/) by Alex Edwards

## My twist
- I'll use SQLite as the DB instead of PostgreSQL (which is the author's choice)
- I'll use Go's router (ServeMux) taking advantages of the improvements made on v1.22

## Endpoints
| Method | URL             | Action                                 |
| ------ | --------------- | -------------------------------------- |
| GET    | /v1/healthcheck | Show application information           |
| GET    | /v1/movies      | Show the details of all movies         |
| POST   | /v1/movies      | Create a new movie                     |
| GET    | /v1/movies/:id  | show the details of a specific movie   |
| PATCH  | /v1/movies/:id  | Update the details of a specific movie |
| DELETE | /v1/movies/:id  | Delete a specific movie                |

## DB 
- The application requires a sqlite3 DB file.
- The application expects a `GREENLIGHT_DB_DSN` environment variable, startup will fail if none provided.
- The number of connections of the DB pool must be `1` to minimize the chances of DB locks
  - This is a limitation of goroutines + sqlite3, as sqlite3 doesn't support multiple concurrent writers

## Migrations
To run a migration:
```shell
migrate -path=./migrations -database <DB DSN> 
```
Obs: don't use query strings on DSN!
