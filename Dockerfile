FROM debian:stretch
MAINTAINER Ronoaldo JLP <ronoaldo@gmail.com>

RUN DEBIAN_FRONTEND=non-interactive \
	sed -e 's/deb.debian.org/ftp.us.debian.org/g' -i /etc/apt/sources.list && \
       	apt-get update &&\
	apt-get install -y xvfb phantomjs ca-certificates

ARG GIT_HASH
ENV VERSION GIT_HASH

ADD render.js  /var/lib/
ADD pagerender /usr/bin/

EXPOSE 8080
ENTRYPOINT pagerender
