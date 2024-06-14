build:
	@go build -o bin/matcherx cmd/main.go	

run: build
	@./bin/matcherx

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

# run:
# 	@go run main.go auth.go api.go config.go db.go users.go posts.go store.go utils.go types.go

# test:
# 	@go test -v ./..