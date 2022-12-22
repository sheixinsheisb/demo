FROM debian:11-slim

ADD listBucket /bin

RUN apt update
RUN apt install -y ca-certificates

ENTRYPOINT /bin/listBucket