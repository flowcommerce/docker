# docker
Central place for dockerfiles and other utilities to support docker at flow

## play

[templates/play/Dockerfile](templates/play/Dockerfile) defines a base image that we can use when building our
play applications.

To create a new play base image

    cd play
    go run build.go

This will create a new image using the git tag from this repository,
build the image, and tag w/ latest.


## postgresql

templates/postgresql/Dockerfile defines a base image that we can use when building our
postgresql applications.

    FROM flowdocker/postgresql:0.0.71

To create a new postgresql base image

    cd postgresql
    go run $GOPATH/src/github.com/flowcommerce/tools/dev.go tag
    go run build.go


## node
See [node README.md](node/README.md) for details.
