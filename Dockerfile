FROM golang:1.16-alpine

WORKDIR $GOPATH/src/github.com/kilianstallt/mongo-snap

RUN apk add mongodb-tools

COPY . .

RUN go mod download

RUN go install .

WORKDIR /

ENTRYPOINT [ "mongo-snap" ]
