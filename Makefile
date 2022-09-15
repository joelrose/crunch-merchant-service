start-services:
	docker compose -f .development/docker-compose.yml up -d 
run:
	go run main.go
gen-swagger:
	swag init