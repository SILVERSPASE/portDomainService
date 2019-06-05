FROM golang:1.12.2

ENV GO111MODULE=on

WORKDIR /go/src/github.com/silverspase/portService/clientAPI/
RUN mkdir -p /go/src/github.com/silverspase/portService/clientAPI/
COPY . /go/src/github.com/silverspase/portService/clientAPI/
RUN go mod vendor
RUN GOOS=linux GOARCH=amd64 go build ./cmd/main.go

CMD ["./main", "-server=server:9090", "-port=:8080"]