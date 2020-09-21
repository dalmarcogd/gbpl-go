.PHONY: protos

protos:
	protoc -I=./ --go_out=./  grpcs.proto
	protoc -I=./ --go-grpc_out=./ grpcs.proto
	mv github.com/dalmarcogd/gbpl-go/pkg/grpcs/grpcs.pb.go ../../pkg/grpcs/grpcs.pb.go
	mv github.com/dalmarcogd/gbpl-go/pkg/grpcs/grpcs_grpc.pb.go ../../pkg/grpcs/grpcs_grpc.go
	rm -rf github.com
