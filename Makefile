start-services:
	docker compose -f .development/docker-compose.yml up -d 
run:
	go run main.go
gen-swagger:
	swag init
stripe-webhook:
	stripe listen --events charge.succeeded --forward-to localhost:8080/api/v1/webhook/stripe