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

- add a .env file on the root folder with the DB config variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_MAX_IDLE_CONNS, DB_MAX_OPEN_CONNS, TOKEN, REFRESHTOKEN, SECRET_KEY) or fill out the given one with the correct DB credentials and config values (you can copy an rename .env.example)
- run `make lint`
- run `make dev`
- run `make test-unit`

## Run e2e tests

- with a local server instance listening, run:
- `make test-e2e`
- **Note**: This will actually execute the endpoints under tests, so make sure your env file is pointing to the right environment

## Deployment in render

### DB

In Render select "New" and **PostgreSQL**. It will show a screen with some pre-configuration that can be done regarding naming, default user and versioning. Click on "Create DB" and you are done.

Once created it will lead you to a screen where you can monitor and check the connection info of the DB. You can connect to it from `PGAdmin` to run the migrations needed.

### Service

In Render select "New" and **web service**, as we are using a public repository, set the service to read the git repository.

A config screen will pop where you can set the right values for your own DB, preferences and the environment variables needed("Advanced" button at the bottom) (copy them from .env.example). Important ones:

> Build command: `go get . && go build -tags netgo -ldflags '-s -w' -o app`

> Health Check Path: `/`
