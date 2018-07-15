FROM ubuntu:16.04

RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD main /main
ADD entrypoint.sh /entrypoint.sh
ADD static /static
WORKDIR /

EXPOSE 8050
ENTRYPOINT ["/entrypoint.sh"]