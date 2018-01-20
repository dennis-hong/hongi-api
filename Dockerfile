FROM golang:1.8.3

ENV SRC_DIR $GOPATH/src/github.com/dennis-hong/hongi-api

COPY . $SRC_DIR

WORKDIR $SRC_DIR
RUN go get
RUN go build -o %GOPATH/bin/hongi

CMD $GOPATH/bin/hongi