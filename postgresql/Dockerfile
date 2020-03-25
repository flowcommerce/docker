FROM ubuntu:16.04

# This is a common docker file for flow that is used to create
# containers running specific versions of an applications schema. This
# container will be running postgresql with the application's schema
# applied.

# We make an assumption that there is a file named 'install.sh' in the
# same directory as this Dockerfile that contains the instructions for
# creating the application database, user, etc.

# Each schema is intended to be in its own git repository so that we
# can link that repository to docker hub to automatically build the
# docker images following a tag of the repository.

# Example from command line:
#
#  docker build -t flowcommerce/splashpage-postgresql:0.1.2 .
#
#  docker run -d -p 5100:5432 flowcommerce/splashpage-postgresql:0.1.2
#
#  psql -U api -h 192.168.99.100 -p 5100 splashpagedb
#

MAINTAINER tech@flow.io

RUN apt-get update
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/postgresql.list
RUN apt-get install -y --no-install-recommends ca-certificates wget sudo
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

RUN apt-get update
RUN apt-get install -y postgresql-11
RUN apt-get install -y --no-install-recommends git
RUN apt-get install -y --no-install-recommends ruby
RUN apt-get install -y --no-install-recommends curl

WORKDIR /opt
RUN git clone git://github.com/mbryzek/schema-evolution-manager.git
WORKDIR /opt/schema-evolution-manager
RUN ls
RUN git checkout 0.9.35
RUN ruby ./configure.rb --prefix /usr/local
RUN ./install.rb

RUN sed -i 's/peer/trust/' /etc/postgresql/11/main/pg_hba.conf
RUN sed -i 's/md5/trust/' /etc/postgresql/11/main/pg_hba.conf
RUN sed -i 's/127.0.0.1\/32/0.0.0.0\/0/' /etc/postgresql/11/main/pg_hba.conf
RUN cat /etc/postgresql/11/main/pg_hba.conf

RUN cat "/etc/init.d/postgresql"

RUN mkdir -p /var/run/postgresql && chown -R postgres /var/run/postgresql

WORKDIR /var/lib/postgresql/11/main
RUN ln -s /etc/postgresql/11/main/postgresql.conf

RUN sed -i 's/include_dir/#include_dir/' /etc/postgresql/11/main/postgresql.conf

VOLUME /var/lib/postgresql/11/base
EXPOSE 5432

# see - https://github.com/nimiq/docker-postgresql93/pull/4/files
RUN mkdir /etc/ssl/private-copy; mv /etc/ssl/private/* /etc/ssl/private-copy/; rm -r /etc/ssl/private; mv /etc/ssl/private-copy /etc/ssl/private; chmod -R 0700 /etc/ssl/private; chown -R postgres /etc/ssl/private
