.PHONY: protos

protos:
	protoc -I=./idl --go_out=plugins=grpc:./gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/common/types.proto
	protoc -I=./idl --go_out=plugins=grpc:./gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/application/application.proto
#	protoc -I=./idl --go_out=plugins=grpc:./gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/lifecycle/lifecycle.proto
#	protoc -I=./idl --go_out=plugins=grpc:./gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/listings/listings.proto
#	protoc -I=./idl --go_out=plugins=grpc:./gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/users/users.proto

