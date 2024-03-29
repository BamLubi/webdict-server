# webdict-server

![](https://img.shields.io/badge/-Go-brightgreen) ![](https://img.shields.io/badge/-Docker-blue) ![](https://img.shields.io/badge/-CI/CD-red)

Webdict-Server is a background service, builted by Go, and used to support WebDict. It has completed the CI/CD, which will automatic build and deploy the project when code changed.

## Environment

- Go 1.19
- Docker 20.10.19
- Mysql 5.7
- Redis latest

## Quick Start

1. Setup Docker Environment. If you are not familiar with Docker, Please read [Docker Installation](#Docker_Installation).

```sh
docker-compose -f docker-compose.yml up -d
```

1. Run Project.

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

4. (optional) Import allowed hosts to the redis server.

```shell
python gen_allowed_host.py
```

<a id="Docker_Installation"></a>
## Docker Installation

1. Delete all old version about docker.

```sh
apt-get remove docker docker-engine docker.io
```

2. Install require packages.

```sh
apt-get install apt-transport-https ca-certificates curl software-properties-common
```

3. Configure the download settings.

```sh
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
```

4. Install docker and docker-compose.

```sh
apt-get update
apt-get install docker-ce docker-compose
```

5. Run Docker.

```sh
systemctl enable docker
systemctl start docker
```

## CI/CD
supproted by `GitHub Actions`

When you push code to the branch "main", it will trigger the Action "Build & Deploy", which will build the project and deploy it to the remote server(you need to config your own secrets).

Please check the file `.github\workflows` for more details.
