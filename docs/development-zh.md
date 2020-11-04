# 开发

## 环境准备
### MinIO
```　bash
docker run -p 9000:9000 -d --name=minio -e MINIO_ACCESS_KEY=minio -e MINIO_SECRET_KEY=miniostorage minio/minio server /data
```

## 编码
修改了 Repository 之后，需要运行 `make mockgen` 来生成 Mock

### 编译 
`make build`

### lint
使用 golang-ci [golangci-lint](https://github.com/golangci/golangci-lint#editor-integration) 来检查

### 创建数据库
PS：不建议在 Windows 下使用 flyway，坑太多，建议在 wsl 中操作

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

### 创建用户
运行命令：`go run cmd/jellyfish-tool/main.go create-user jellyfish 123456`


## postgres permission
postgres user must be supperuser
`ALTER USER jellyfish WITH SUPERUSER;`      