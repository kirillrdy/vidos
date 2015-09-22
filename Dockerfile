FROM ubuntu:15.10

RUN apt-get update
RUN apt-get install -y postgresql
RUN apt-get install -y vim
RUN apt-get install -y ffmpeg
RUN apt-get install -y git golang
RUN apt-get install -y golang
RUN apt-get install -y vim

ADD install.sh /install.sh
RUN /bin/sh /install.sh

ENV GOPATH /go

ADD . /go/src/github.com/kirillrdy/vidos
RUN go get github.com/kirillrdy/vidos
RUN go install github.com/kirillrdy/vidos


CMD service postgresql start && sleep 3 && /go/bin/vidos
EXPOSE 3001
