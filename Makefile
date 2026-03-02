include .env
export

service-deploy:
	docker compose up -d application

service-undeploy:
	docker compose down application

service-run:
	@go run main.go
migrate-up:
	migrate -path migrations -database ${CONN_STRING} up
migrate-down:
	migrate -path migrations -database ${CONN_STRING} down
