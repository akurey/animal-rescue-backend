## Getting started

- brew install Go@1.18
- brew install golangci-lint         
- `cd YOUR_PROJECT_PATH`
- go get . 


## Project structure

- `Controllers` : keep all your API controllers here
- `database` : keep all your database and queries here
- `models` : keep all your data structs here
- `Test` : keep all your test files here
- `main.go` : main project file


## Run it locally

- add a .env file on the root folder with the DB config variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_MAX_IDLE_CONNS, DB_MAX_OPEN_CONNS) or fill out the given one with the correct DB credentials and config values
- run `golangci-lint run`
- run `go run main.go`
- run `go test ./tests/unit`
