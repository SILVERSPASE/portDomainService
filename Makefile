proto:
	protoc --proto_path=./protorepo/v1 --go_out=plugins=grpc:portDomainService/internal/api/v1 \
	--go_out=plugins=grpc:clientAPI/internal/api/v1 port-service.proto

client:
	cd ./clientAPI && go run cmd/main.go -server=localhost:9090 -port=:8080

backend:
	cd ./portDomainService && go run ./cmd/main.go -grpc-port=9090 -mongoURI=mongodb://localhost:27017 -dbName=domainService

mongo:
	docker run --name mongo -p 27017:27017 -d mongo:latest

json:
	curl localhost:8080/json -X POST -F 'uploadFile=@/Users/ashch/go/src/github.com/silverspase/portService/ports.json'