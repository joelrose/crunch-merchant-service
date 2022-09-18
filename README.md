# Crunch Backend

## How to Run

1. Install Go
2. Start DB with `docker run --name crunch-postgres -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=catalog-service -d postgres:12`
3. Run `copy .env.example .env`
4. Run `go run main.go`

## Database

### Setup

1. Install `brew install golang-migrate`
2. Start postgres docker locally: `docker run --name merchant-service-postgres -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=catalog-service -d postgres:12`
3. `cp .env.example .env` and `source .env`

Execute migrations in the db:

`migrate -path migrations -database $DATABASE_URL -verbose up`

Roll the migrations back:

`migrate -path migrations -database $DATABASE_URL -verbose down`

Create a migration:

`migrate create -ext sql -dir migrations -seq name`

## Authentication

### Merchant Dashboard: Auth0

We use Auth0 for authentication and authorization in the crunch dashboard. Protected routes have to use the middleware Auth0Auth. 

### Consumer App: Firebase

We use Firebase for the authentication in the crunch app. Protected routes have to use the middleware FirebaseAuth. 

## Redis

Redis is used to cache the store menus.