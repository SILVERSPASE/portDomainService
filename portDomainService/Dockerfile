FROM golang:1.12.2

ENV GO111MODULE=on

WORKDIR /go/src/github.com/silverspase/portService/portDomainService/
RUN mkdir -p /go/src/github.com/silverspase/portService/portDomainService/
COPY . /go/src/github.com/silverspase/portService/portDomainService/
RUN go mod vendor
RUN GOOS=linux GOARCH=amd64 go build ./cmd/main.go

CMD ["./main", "-grpc-port=9090", "-mongoURI=mongodb://mongodb:27017", "-dbName=domainService"]