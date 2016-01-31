# docker
Central place for dockerfiles and other utilities to support docker at flow

## play

play/Dockerfile defines a base image that we can use when building our
play applications.

    FROM flowcommerce/play:0.0.8

If you modify the Dockerfile and want to create a new image:

    /web/tools/bin/tag
    cd play
    ./build

This will create a new image using the git tag from this repository,
build the image, tag w/ latest, and update this README so the example
is the latest version of the play image.