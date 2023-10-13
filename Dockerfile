FROM golang:1.21.1-alpine

RUN mkdir -p /go/src/love-comme-server/
COPY . /go/src/love-comme-server/
WORKDIR /go/src/love-comme-server/

RUN go install

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air"]
