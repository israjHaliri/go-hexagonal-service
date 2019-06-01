## Needed /Tech stacks
    + mysql
    + go 1.11 above (follow this for installation : https://golang.org/doc/install)

## To get started follow this checklist:
    + create schema depends on datasource.go
    + run migration go run cmd/migration/main.go
    + get into directory project then type : cd cmd && go run main.go

## To build follow this checklist:
    + get into directory project then type : cd cmd && go build -o {app_name}

## To run testing
    + run migration go run cmd/migration/main.go
    + go test -v ./pkg/listing ./pkg/saving/ ./pkg/storage/database/ -count=1 -coverprofile coverage.out && go tool cover -html=coverage.out

## Note: 
    ++ for production use supervisor using commad go run {app_name}
    ++ pattern take from https://github.com/sourcegraph/about/blob/master/blogposts/gophercon-2018-how-do-you-structure-your-go-apps.md
