FROM golang:1.23.3-alpine3.20 as base
RUN apk add --no-cache curl

WORKDIR /usr/src/app
COPY . .

RUN chmod +x /usr/src/app/docker/scripts/start.sh

CMD ["/bin/sh", "/usr/src/app/docker/scripts/start.sh"]