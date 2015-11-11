FROM ubuntu:15.10

RUN apt-get update && apt-get install -y postgresql ffmpeg git golang

ADD install.sh /install.sh
RUN /bin/sh /install.sh

ENV GOPATH /go

WORKDIR /go/src/github.com/kirillrdy/vidos

ADD . /go/src/github.com/kirillrdy/vidos


RUN go get -v
RUN go install


CMD service postgresql start && sleep 3 && /go/bin/vidos
EXPOSE 3001
