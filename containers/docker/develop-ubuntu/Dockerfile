FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/smartshareproject/smartshare && \
  (cd smartshare && make smartshare) && \
  cp smartshare/build/bin/smartshare /smartshare && \
  apt-get remove -y golang-1.9 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /smartshare

EXPOSE 8545
EXPOSE 18883

ENTRYPOINT ["/smartshare"]
