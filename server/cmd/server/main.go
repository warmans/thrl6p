package main

import (
	_ "github.com/lib/pq"
	"github.com/warmans/thrl6p/server/pkg/server"
	"github.com/warmans/thrl6p/server/pkg/service/grpc"
	"github.com/warmans/thrl6p/server/pkg/service/http"
	"github.com/warmans/thrl6p/server/pkg/store"
	"go.uber.org/zap"
	"os"
)

func main() {

	logger := newLogger()
	defer logger.Sync()

	db, err := store.NewConn(os.Getenv("DB_DSN")) // todo flags
	if err != nil {
		logger.Fatal("failed to create DB connection", zap.Error(err))
	}
	logger.Info("Running DB migrations...")
	if err := db.Migrate(); err != nil {
		logger.Fatal("failed to migrate DB", zap.Error(err))
	}

	grpcServices := []server.GRPCService{
		grpc.NewPatchService(grpc.PatchServiceConfig{}, db),
	}
	httpServices := []server.HTTPService{
		http.NewDownloadService(),
	}

	// todo: flags
	conf := server.GrpcServerConfig{
		GRPCAddr: "0.0.0.0:9090",
		HTTPAddr: ":8888",
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
