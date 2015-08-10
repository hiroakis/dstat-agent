FROM ubuntu

MAINTAINER Hiroaki Sano <hiroaki.sano.9stories@gmail.com>

RUN apt-get update -y && apt-get install git golang dstat -y
RUN mkdir /tmp/go \
    && export GOPATH=/tmp/go \
    && go get github.com/hiroakis/dstat-agent

ENV PATH /tmp/go/bin:$PATH

EXPOSE 8888

CMD ["dstat-agent"]
