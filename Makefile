start-services:
	docker compose -f .development/docker-compose.yml up -d 
run:
	go run main.go
stripe-webhook:
	stripe listen --events charge.succeeded --forward-to localhost:8080/api/v1/webhook/stripe
gen-swagger:
	swag init --parseDependency
gen-mocks:
	mockgen -source=db/interface.go -destination=test_helper/mock_db/interface.go && mockgen -source=services/http_client/http_client.go -destination=test_helper/mock_http_client/interface.go
test:
	go test -v -cover ./...
db-import:
	PGPASSWORD=password psql -h localhost -p 5432 -U username database -f .development/db_dump.sql
