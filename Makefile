start-services:
	docker compose -f .development/docker-compose.yml up -d 
run:
	go run main.go
stripe-webhook:
	stripe listen --events charge.succeeded --forward-to localhost:8080/api/v1/webhook/stripe
gen-swagger:
	swag init
gen-mocks:
	mockgen -source=db/interface.go -destination=test_helper/mock_db/interface.go
test:
	go test -cover ./...