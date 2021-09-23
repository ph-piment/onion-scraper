FROM golang:1.17.0-alpine

RUN apk update
RUN apk add curl

# タイムゾーン設定
ENV TZ Asia/Tokyo

ENV WORKDIR_PATH /go/src/github.com/ph-piment/onion-scraper
WORKDIR ${WORKDIR_PATH}
ADD . ${WORKDIR_PATH}

RUN curl -L https://github.com/k0kubun/sqldef/releases/download/v0.11.4/psqldef_linux_arm64.tar.gz | tar xvz -C /usr/local/bin
