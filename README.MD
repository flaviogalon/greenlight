# Greenlight
JSON API for retrieving and managing information about movies.

This is an educational project developed following the book [Let's Go Further](https://lets-go-further.alexedwards.net/) by Alex Edwards

## My twist
- I'll use SQLite as the DB instead of PostgreSQL (which is the author's choice)
- I'll use Go's router (ServeMux) taking advantages of the improvements made on v1.22

## Endpoints
| Method | URL             | Action                               |
| ------ | --------------- | ------------------------------------ |
| GET    | /v1/healthcheck | Show application information         |
| POST   | /v1/movies      | Create a new movie                   |
| GET    | /v1/movies/:id  | show the details of a specific movie |
