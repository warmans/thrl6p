// +build tools

package server

import (
	_ "github.com/gogo/protobuf/protoc-gen-gofast"
	_ "github.com/go-bindata/go-bindata/go-bindata"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	//_ "github.com/uber/prototool/cmd/prototool" // causes tidy to fail for some reason
)