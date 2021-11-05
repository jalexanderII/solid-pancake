.PHONY: application_protos

application_protos:
	protoc -I=./services/application/protos --go_opt=paths=source_relative --go_out=plugins=grpc:./services/application/protos/ ./services/application/protos/application.proto