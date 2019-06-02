proto:
	cd ./portDomainService && protoc --proto_path=api/v1 --go_out=plugins=grpc:internal/api/v1 port-service.proto && \
	cp ./internal/api/v1/port-service.pb.go ../clientAPI/internal/api/v1/port-service.pb.go

client:
	cd ./clientAPI && go run cmd/main.go -server=localhost:9090 -port=:8080