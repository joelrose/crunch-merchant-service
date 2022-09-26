start-services:
	docker compose -f .development/docker-compose.yml up -d 
run:
	go run main.go
stripe-webhook:
	stripe listen --events charge.succeeded --forward-to localhost:8080/api/v1/webhook/stripe
gen-swagger:
	swag init --parseDependency
gen-mocks:
	mockgen -source=db/interface.go -destination=test_helper/mock_db/mock_db.go && mockgen -source=services/http_client/http_client.go -destination=test_helper/mock_http_client/mock_http_client.go
test:
	go test -cover ./...
db-import:
	PGPASSWORD=password psql -h localhost -p 5432 -U username database -f .development/db_dump.sql
