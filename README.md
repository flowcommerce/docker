# docker
Central place for dockerfiles and other utilities to support docker at flow

## play

templates/play/Dockerfile defines a base image that we can use when building our
play applications.

    FROM flowdocker/play:0.0.16

To create a new play base image

    /web/tools/bin/tag
    cd play
    ./build

This will create a new image using the git tag from this repository,
build the image, tag w/ latest, and update this README so the example
is the latest version of the play image.


## postgresql

templates/postgresql/Dockerfile defines a base image that we can use when building our
postgresql applications.

    FROM flowdocker/postgresql:0.0.9

To create a new postgresql base image

    cd postgresql
    go build.go
