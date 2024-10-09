docker-up:
	docker compose up -d

docker-down:
	docker compose down

goose-up:
	goose -dir migrations postgres "user=stockradar password=stockradar dbname=stockradar port=5433 sslmode=disable" up

goose-down:
	goose -dir migrations postgres "user=stockradar password=stockradar dbname=stockradar port=5433 sslmode=disable" down

server:
	go run main.go