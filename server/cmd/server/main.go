package main

import (
	_ "github.com/lib/pq"
	"github.com/warmans/thrl6p/server/pkg/flag"
	"github.com/warmans/thrl6p/server/pkg/server"
	"github.com/warmans/thrl6p/server/pkg/service/grpc"
	"github.com/warmans/thrl6p/server/pkg/service/http"
	"github.com/warmans/thrl6p/server/pkg/store"
	"go.uber.org/zap"
)

const ServicePrefix = "THRL6P"

func main() {

	dbCfg := &store.Config{}
	dbCfg.RegisterFlags(ServicePrefix)

	grpcCfg := server.GrpcServerConfig{}
	grpcCfg.RegisterFlags(ServicePrefix)

	flag.Parse()

	logger := newLogger()
	defer logger.Sync()

	db, err := store.NewConn(dbCfg)
	if err != nil {
		logger.Fatal("failed to create DB connection", zap.Error(err))
	}
	logger.Info("Running DB migrations...")
	if err := db.Migrate(); err != nil {
		logger.Fatal("failed to migrate DB", zap.Error(err))
	}

	grpcServices := []server.GRPCService{
		grpc.NewPatchService(db),
	}
	httpServices := []server.HTTPService{
		http.NewDownloadService(),
	}

	srv, err := server.NewServer(logger, grpcCfg, grpcServices, httpServices)
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
	// the stack traces are mostly useless because they're usually the trace of the log call.
	logger, err := zap.NewProduction(zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		panic(err)
	}
	return logger
}
