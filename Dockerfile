FROM ubuntu:15.10

RUN apt-get update && apt-get install -y postgresql ffmpeg git golang

ADD install.sh /install.sh
RUN /bin/sh /install.sh

ENV GOPATH /go

ADD . /go/src/github.com/kirillrdy/vidos
RUN go get github.com/kirillrdy/vidos
RUN go install github.com/kirillrdy/vidos


CMD service postgresql start && sleep 3 && /go/bin/vidos
EXPOSE 3001
