FROM node:$NODE_VERSION-alpine

LABEL maintainer="tech@flow.io"

# Default to UTF-8 file.encoding
ENV LANG C.UTF-8

ENV NODE_ENV production

WORKDIR /root

COPY environment-provider.jar .
COPY environment-provider-version.txt .

RUN apk update && apk upgrade \
&& apk add --no-cache openjdk11-jre \
&& npm install -g forever \
&& apk add --no-cache --update ca-certificates curl \
&& mkdir -p /opt/node/log

COPY .npmrc /opt/node/.npmrc

WORKDIR /opt/node
