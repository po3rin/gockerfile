# gockerfile

<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/> [![GolangCI](https://golangci.com/badges/github.com/po3rin/gockerfile.svg)](https://golangci.com) <a href="https://codeclimate.com/github/po3rin/gockerfile/maintainability"> <a href="https://codeclimate.com/github/po3rin/gockerfile/maintainability"><img src="https://api.codeclimate.com/v1/badges/7cc6dbab602cfd7e2e9a/maintainability" /></a>

:whale:
gockerfile is a YAML Docker-compatible alternative to the Dockerfile Specializing in simple go server.

## Instalation as cmd

```bash
$ go get -u github.com/po3rin/gockerfile/cmd/gocker
```

## Usage

### po3rin/gocker config file

create Gockerfile.yaml (Gockerfile supports only 3 fields)

```yaml
#syntax=po3rin/gocker

repo: github.com/po3rin/gockerfile
path: ./example/server
version: v0.0.1 # default master
```

repo is git repository. path is path to main.go.

### Build Gockerfile using docker build

you can build Gockerfile.yaml with docker build

```
$ DOCKER_BUILDKIT=1 docker build -f Gockerfile.yaml -t po3rin/gockersample .
```

### Build Gockerfile with builtctl

using as buildkit frontend.

```bash
buildctl build \
		--frontend=gateway.v0 \
		--opt source=po3rin/gocker \
		--local gockerfile=. \
		--output type=docker,name=gockersample | docker load
```

## Run container

You can run go API container as it is.

```bash
$ docker run -it -p 8080:8080 po3rin/gockersample:latest /bin/server
$ curl localhost:8080/
Hello World
```
