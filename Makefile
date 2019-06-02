proto:
	cd ./portDomainService && protoc --proto_path=api/v1 --go_out=plugins=grpc:internal/api/v1 port-service.proto && \
	cp ./internal/api/v1/port-service.pb.go ../clientAPI/internal/api/v1/port-service.pb.go

client:
	cd ./clientAPI && go run cmd/main.go -server=localhost:9090 -port=:8080

service:
	cd ./portDomainService && go run ./cmd/main.go -grpc-port=9090 -mongoURI=mongodb://localhost:27017 -dbName=domainService

mongo:
	docker run --name mongo -p 27017:27017 -d mongo:latest