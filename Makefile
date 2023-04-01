run:
	go run main.go

migration:
	migrate create -ext sql -dir model/migration -seq $(name)

migrate-up:
	migrate -path models/migrations -database "postgres://postgres:kadatahu@localhost:5432/graphql?sslmode=disable" up

migrate-down:
	migrate -path models/migrations -database "postgres://postgres:kadatahu@localhost:5432/graphql?sslmode=disable" down