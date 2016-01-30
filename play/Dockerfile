FROM ubuntu:15.10

MAINTAINER tech@flow.io

# need to run to make sure the apt-get installs work
RUN apt-get update

RUN apt-get install -y --no-install-recommends ca-certificates wget
RUN apt-get install -y --no-install-recommends ca-certificates software-properties-common
RUN echo debconf shared/accepted-oracle-license-v1-1 select true | debconf-set-selections
RUN echo debconf shared/accepted-oracle-license-v1-1 seen true | debconf-set-selections
RUN add-apt-repository -y ppa:webupd8team/java
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates oracle-java8-installer
RUN javac -version # test

RUN echo "deb https://dl.bintray.com/sbt/debian /" | tee -a /etc/apt/sources.list.d/sbt.list
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 642AC823
RUN apt-get install apt-transport-https
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates sbt=0.13.9

# add .sbt and .ivy2 to root directory
WORKDIR /root
RUN mkdir .sbt/
RUN mkdir .ivy2/
ADD .sbt/ .sbt/.
ADD .ivy2 .ivy2/.

ADD environment-provider.jar .
ADD environment-provider-version.txt .
