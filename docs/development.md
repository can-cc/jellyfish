# Development

### compile 
`go build main.go`

### lint
this project use [golangci-lint](https://github.com/golangci/golangci-lint#editor-integration)

### create database
docker:
``` bash
docker run --rm -v ./migration/sql:/flyway/sql -v ./migration/conf:/flyway/conf boxfuse/flyway migrate 
```

docker command line:
``` bash
docker run --rm -v $(pwd)/migration/sql:/flyway/sql boxfuse/flyway migrate -url=jdbc:postgresql://172.17.0.1:5432/jellyfish -user=postgres -password=mysecretpassword
```

command line:
``` bash
flyway migrate -url=jdbc:postgresql://localhost:5432/jellyfish -user=postgres -password=mysecretpassword -locations="./migration/sql"
```

### create user
`./create-user fwchen 123456`
