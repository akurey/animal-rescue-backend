## Getting started

- brew install Go@1.18
- `cd YOUR_PROJECT_PATH`
- Setup your .env file, for example follow .env.example
- To add all dependencies for a package in your module `go get .` in the current directory


## Project structure

- `Controllers` : keep all your API controllers here
- `repo` : keep all your database and queries here
- `models` : keep all your data structs here
- `middleware` : keep all the middleware files here
- `routes` : keep all routing files here
- `main.go` : main project file

## Run it locally
- `go run main.go` or `go build main.go` and run `./main`
- The application should be available and running on localhost:8000
