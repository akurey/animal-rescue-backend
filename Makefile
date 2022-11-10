dev:
	go run main.go

lint:
	golangci-lint run

test-unit:
	go test ./tests/unit

test-e2e:
	APIURL=http://localhost:8080 ./tests/e2e/run-api-tests.sh