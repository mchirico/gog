[![Build Status](https://travis-ci.org/mchirico/gog.svg?branch=develop)](https://travis-ci.org/mchirico/gog) [![Maintainability](https://api.codeclimate.com/v1/badges/1558bc5ede187bd55266/maintainability)](https://codeclimate.com/github/mchirico/gog/maintainability)

# gog
Go Gorilla Mux -- My Template


## Install

```bash
# Dependencies
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/rpc

#  ... client: (grpc_client)
go get -u github.com/levigross/grequests


# Run Different Examples
go get -u github.com/mchirico/gog/cmd/gog
go get -u github.com/mchirico/gog/cmd/grpc
```



# Create Minimal Docker Image

Reference [minimal](https://github.com/mchirico/gog/tree/min_docker/docker/minimal)


## Dockerfile
```bash
FROM       scratch

MAINTAINER Mike Chirico (chico) <mchirico@gmail.com>
ADD        gog gog

```

## Commands

```bash
#!/bin/bash
env GOOS=linux GOARCH=arm go build github.com/mchirico/gog/cmd/gog
docker build --tag gogserver .
echo -e '\n\n\nIn another window execute the following command:\n\ncurl localhost:8080\n\n'
docker run -p 8080:8080 --rm gogserver

```


## Listing of Example Programs
[awesome-go](https://awesome-go.com/)