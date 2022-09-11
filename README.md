# Crunch Backend

## How to Run

1. Install Go
2. Start DB with `docker run --name merchant-service-postgres -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=catalog-service -d postgres:12-alpine`
3. Run `go run main.go`

## Database

### Setup

1. Install `brew install golang-migrate`
2. Start postgres docker locally: `docker run --name merchant-service-postgres -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=catalog-service -d postgres:12-alpine`
3. `cp .env.example .env` and `source .env`

Execute migrations in the db:

`migrate -path migrations -database $DATABASE_URL -verbose up`

Roll the migrations back:

`migrate -path migrations -database $DATABASE_URL -verbose down`

Create a migration:

`migrate create -ext sql -dir migrations -seq name`
