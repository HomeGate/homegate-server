FROM golang

ADD . /go/src/homegate-server
WORKDIR /go/src/homegate-server

RUN go get -u github.com/golang/dep/...
RUN dep ensure
RUN go install

EXPOSE 8080

ENTRYPOINT /go/bin/homegate-server
