.PHONY: protos

protos:
	protoc -I=./idl --go_opt=paths=source_relative --go_out=plugins=grpc:./gen ./idl/common/types.proto --go-grpc_out=require_unimplemented_servers=false:gen
	protoc -I=./idl --go_opt=paths=source_relative --go_out=plugins=grpc:./gen ./idl/application/application.proto --go-grpc_out=require_unimplemented_servers=false:gen
#	protoc -I=./idl --go_opt=paths=source_relative --go_out=plugins=grpc:./gen ./idl/lifecycle/lifecycle.proto --go-grpc_out=require_unimplemented_servers=false:gen
#	protoc -I=./idl --go_opt=paths=source_relative --go_out=plugins=grpc:./gen ./idl/listings/listings.proto --go-grpc_out=require_unimplemented_servers=false:gen
#	protoc -I=./idl --go_opt=paths=source_relative --go_out=plugins=grpc:./gen ./idl/users/users.proto --go-grpc_out=require_unimplemented_servers=false:gen

