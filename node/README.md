# flowdocker/node* docker images

## Overview

These node images provide a standard node environment for use by Flow's Nodejs apps. This includes
the required auth for accessing Flow's private modules in the npm registry, and Flow's
environment_providor.jar to give access to the appropriate application environment at runtime.
There are 2 flavours of the node image. The basee 'node' image and the 'node_builder' image.
See below for details.

### 'flowdocker/node*' image

This image is intended to be used by node apps in production. I contains the basic setup required to
run a Flow node app.

#### build

```
./build-node <docker_image_version_tag> <node_major_version>
```

### 'flowdocker/node*_builder' image

This image effectively extends the `flowdocker/node` image with extra tools for building frontend assets
and uploading them to the CDN etc.. It includes the Flow 'dev' tool to facilitate this.

#### build

requirements
 * `GITHUB_TOKEN` environment variable with a valid [Github personal access token](https://github.com/settings/tokens) with repo access. This is required to allow `dep ensure` access to Flow's private repo's during build.

```
./build-node_docker <docker_image_version_tag> <node_major_version>
```