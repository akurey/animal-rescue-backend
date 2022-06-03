## Getting started

- brew install Go@1.18
- brew install golangci-lint         
- `cd YOUR_PROJECT_PATH`
- go get . 


## Project structure

- `Controllers` : keep all your API controllers here
- `repo` : keep all your database and queries here
- `models` : keep all your data structs here
- `Test` : keep all your test files here
- `main.go` : main project file


## Run it locally

- run `golangci-lint run`
- run `go run main.go`
- run `go test ./tests/unit`
