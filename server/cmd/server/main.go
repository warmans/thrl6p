package main

import (
	"github.com/warmans/thrl6p/server/pkg/server"
	"github.com/warmans/thrl6p/server/pkg/server/grpc"
	"github.com/warmans/thrl6p/server/pkg/server/http"
	"go.uber.org/zap"
)

func main() {

	logger := newLogger()
	defer logger.Sync()

	conf := server.GrpcServerConfig{
		GRPCAddr: "0.0.0.0:9090",
		HTTPAddr: ":8888",
	}
	grpcServices := []server.GRPCService{
		grpc.NewPatchService(grpc.PatchServiceConfig{}),
	}
	httpServices := []server.HTTPService{
		http.NewDownloadService(),
	}

	srv, err := server.NewServer(logger, conf, grpcServices, httpServices)
	if err != nil {
		logger.Fatal("failed to create server", zap.Error(err))
	}
	go func() {
		if err := srv.StartGRPC(); err != nil {
			logger.Fatal("GRPC Failed", zap.Error(err))
		}
	}()
	if err := srv.StartHTTP(); err != nil {
		logger.Fatal("HTTP Failed", zap.Error(err))
	}
}

func newLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

