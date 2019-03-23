# gockerfile

<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/> [![GolangCI](https://golangci.com/badges/github.com/po3rin/gockerfile.svg)](https://golangci.com) <a href="https://codeclimate.com/github/po3rin/gockerfile/maintainability"> <a href="https://codeclimate.com/github/po3rin/gockerfile/maintainability"><img src="https://api.codeclimate.com/v1/badges/7cc6dbab602cfd7e2e9a/maintainability" /></a>

gockerfile is a YAML Docker-compatible alternative to the Dockerfile Specializing in simple go server.

## Instalation

```bash
$ go get -u github.com/po3rin/gockerfile/cmd/gocker
```

## Usage

create Gockerfile.yaml (Gockerfile supports only simple format)

```yaml
#syntax=po3rin/gocker

repo: github.com/po3rin/gockerfile
path: ./example/server
```

run go api server from repository source code. repo is git repository. path is path to main.go.


gocker lets you build image from Gockerfile using buildctl & docker expoter.

```bash
$ gocker | buildctl build --exporter=docker --exporter-opt name=gockersample | docker load
```

or using as buildkit frontend.

```bash
buildctl build \
		--frontend=gateway.v0 \
		--frontend-opt=source=po3rin/gocker \
		--local dockerfile=. \
		--local context=. \
		--exporter=docker \
		--exporter-opt name=gockersample | docker load
```

you can exec go api container.

```bash
$ docker run -it -p 8080:8080 po3rin/gockersample:latest /bin/server
$ curl localhost:8080/
Hello World
```
