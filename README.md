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

## Creating a local DB

- `docker compose up -d` will create a local postgres instance in port 5438 (make sure your env file is right)


## Run it locally

- add a .env file on the root folder with the DB config variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_MAX_IDLE_CONNS, DB_MAX_OPEN_CONNS) or fill out the given one with the correct DB credentials and config values (you can copy an rename .env.example)
- run `make lint`
- run `make dev`
- run `make test-unit`

## Run e2e tests

- with a local server instance listening, run:
- `make test-e2e`
- **Note**: This will actually execute the endpoints under tests, so make sure your env file is pointing to the right environment
