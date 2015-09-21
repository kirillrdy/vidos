FROM golang

RUN apt-get update
RUN apt-get -y install postgresql vim
RUN service postgresql start && su postgres -c 'createuser -s root'
RUN service postgresql start && createdb vidos

ADD . /go/src/github.com/kirillrdy/vidos
RUN go get github.com/kirillrdy/vidos
RUN go install github.com/kirillrdy/vidos


CMD service postgresql start && /go/bin/vidos
EXPOSE 3001
