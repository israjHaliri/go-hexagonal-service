## Needed /Tech stacks
    + mysql
    + go 1.11 above (follow this for installation : https://golang.org/doc/install)

## To get started follow this checklist:
    + create schema depends on datasource.go
    + get into directory project then type : cd cmd && go run main.go

## To build follow this checklist:
    + get into directory project then type : cd cmd && go build -o {app_name}

## To run testing
    + go test -v ./pkg/storage/database

## Note: 
    ++ for production use supervisor using commad go run {app_name}
    ++ pattern take from https://github.com/sourcegraph/about/blob/master/blogposts/gophercon-2018-how-do-you-structure-your-go-apps.md
