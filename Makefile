proto:
	cd ./portDomainService && protoc --proto_path=api/v1 --go_out=plugins=grpc:internal/api/v1 port-service.proto && \
	cp ./internal/api/v1/port-service.pb.go ../clientAPI/internal/api/v1/port-service.pb.go