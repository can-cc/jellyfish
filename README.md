<img src="https://ae01.alicdn.com/kf/Ufb8b47658198433f8827b13d64e2b55eu.jpg" alt="y5c1eA.png" border="0" height="120" align="right" />

# jellyfish

## run the server

``` bash
go get

go build main.go

./main 
```

then open [http://localhost:8180]

## Development

see [development docs](./docs/development.md)


## Elastic apm configure
To enable elastic apm, set below environment variables.

``` bash
export ELASTIC_APM_SERVICE_NAME=

# Set custom APM Server URL (default: http://localhost:8200)
export ELASTIC_APM_SERVER_URL=

# Use if APM Server requires a token
export ELASTIC_APM_SECRET_TOKEN=

export ELASTIC_APM_ENVIRONMENT="production"
```
See the [documentation](https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html) for advanced configuration.
