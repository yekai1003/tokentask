FROM golang:latest

MAINTAINER yekai "yekai_23@sohu.com"

WORKDIR $GOPATH/src/tokentask

ADD . $GOPATH/src/tokentask

RUN mkdir -p $GOPATH/src/golang.org/x

RUN git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys

RUN go get -u github.com/gin-gonic/gin

RUN go get -u github.com/go-sql-driver/mysql

RUN go build . 

EXPOSE 8080

ENTRYPOINT ["./tokentask"]