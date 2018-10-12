FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/smartshareproject/smartshare && \
  (cd smartshare && make smartshare) && \
  cp smartshare/build/bin/smartshare /smartshare && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /smartshare && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 18883

ENTRYPOINT ["/smartshare"]
