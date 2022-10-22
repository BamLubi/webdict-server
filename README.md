# webdict-server

Webdict-Server is a background service, builted vy Go, and used to support WebDict.

## Environment

- Go 1.19
- Docker 20.10.19
- Mysql 5.7
- Redis latest

## Quick Start

1. Setup Docker Environment

```sh
docker-compose -f docker-compose.yml up -d
```

2. Run Project

```sh
nohup bash ./run.sh > nohup.log 2>&1 &
```

3. (optional) Build and run the project manually.

```sh
# build
go build ./src/main/main.go

# run examples
./webdict_server [Args]
# Args:
# - "": Production Mode, will load the config "app.properties".
# - "-DEV": Development Mode, will load the config "dev.properties".
# - "-PROD_TEST": Production Mode, will load the config "app.properties", 
#   but will not run the server, just used for test.
```
