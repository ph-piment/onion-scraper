FROM golang:1.17.0-alpine

RUN apk update
RUN apk add curl

# タイムゾーン設定
ENV TZ Asia/Tokyo

ENV WORKDIR_PATH /go/src/github.com/ph-piment/onion-scraper
WORKDIR ${WORKDIR_PATH}
ADD . ${WORKDIR_PATH}

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.10.0/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate
