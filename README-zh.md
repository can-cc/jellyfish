# jellyfish server

## 使用
项目用 Makefile 配置了一系列的命令

启动：
``` bash
make run
```

then open [http://localhost:8180]

## 开发

见 [开发文档](./docs/development-zh.md)

## S3/MinIO
``` bash
docker run -p 9000:9000 -d --name=minio -e MINIO_ACCESS_KEY=minio -e MINIO_SECRET_KEY=miniostorage minio/minio server /data
```

## 配置
启动时，必需要配置 pg 数据库和 s3，还有其他配置项需要配置，具体可以修改配置文件[config.yaml](./config/config.yaml)，项目用 [viper](https://github.com/spf13/viper) 来读取配置，也可以通过环境变量来进行配置

例如，设置 pg 数据库，设置 `JFISH_DATASOURCE_RDS_DATABASE_URL` 这个环境变量即可


### Elastic apm 配置
项目集成 elastic apm 功能，如果需要开启，配置下面的环境变量即可

``` bash
export ELASTIC_APM_SERVICE_NAME=

# Set custom APM Server URL (default: http://localhost:8200)
export ELASTIC_APM_SERVER_URL=

# Use if APM Server requires a token
export ELASTIC_APM_SECRET_TOKEN=

export ELASTIC_APM_ENVIRONMENT="production"
```
See the [documentation](https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html) for advanced configuration.
