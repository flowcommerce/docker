FROM flowdocker/play:0.0.8

ADD . /opt/play

WORKDIR /opt/play

RUN sbt clean stage
  
ENTRYPOINT ["java", "-jar", "/root/environment-provider.jar", "run", "play", "APP", "api/target/universal/stage/bin/APP-api"]