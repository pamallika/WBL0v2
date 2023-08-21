run-worker:
		cd cmd && go run ./consumer.go

run-publisher:
		cd cmd && go run ./publisher.go

run-test-db:
		cd internal/repository/database && go test
run-test-cache:
		cd internal/service/cache && go test
