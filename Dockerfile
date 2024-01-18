FROM golang:latest

WORKDIR /app

ADD . /app

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm

ENV GOBIN /go/bin
ENV PATH $GOBIN:$PATH

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
