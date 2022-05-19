FROM golang:1.16-alpine

WORKDIR /opt/metapod

RUN apk update && apk add curl make

ADD . /opt/metapod

ADD docker-entrypoint.sh /

RUN cd /opt/metapod && chmod +x /docker-entrypoint.sh && make build

EXPOSE 8080/tcp
EXPOSE 8080/udp

RUN find . -type f -not -name 'main' -not -name 'index.js' -not -name 'index.tmpl' -not -name 'index.css' -delete

ENTRYPOINT ["/docker-entrypoint.sh"]