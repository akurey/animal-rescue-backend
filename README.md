## Getting started

- brew install Go@1.18
- brew install golangci-lint         
- `cd YOUR_PROJECT_PATH`
- go get . 


## Project structure

- `Controllers` : keep all your API controllers here
- `repo` : keep all your database and queries here
- `models` : keep all your data structs here
- `main.go` : main project file


## Run it locally

- add a .env file on the root folder with the DB config variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
- run `golangci-lint run`
- run `go run main.go`
