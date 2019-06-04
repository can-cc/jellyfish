# jellyfish server


## development
### compile 
`go build main.go`

### create database
docker:
``` bash
docker run --rm -v ./sql:/flyway/sql -v ./conf:/flyway/conf boxfuse/flyway migrate 
```


docker command line:
``` bash
docker run --rm -v ./sql:/flyway/sql boxfuse/flyway migrate -url=jdbc:postgresql://localhost:5432/jellyfish -user=postgres -password=mysecretpassword
```

command line:
``` bash
flyway migrate -url=jdbc:postgresql://localhost:5432/jellyfish -user=postgres -password=mysecretpassword -locations="./sql"
```

### create user
`./create-user fwchen 123456`

