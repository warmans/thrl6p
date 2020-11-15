package grpc

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/teris-io/shortid"
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
	"google.golang.org/grpc"
)

type PatchServiceConfig struct {
}

func NewPatchService(cfg PatchServiceConfig) *PatchService {
	return &PatchService{cfg: cfg}
}

type PatchService struct {
	cfg  PatchServiceConfig
	temp *v1.Patch // dev only
}

func (b *PatchService) RegisterGRPC(server *grpc.Server) {
	v1.RegisterPatchServiceServer(server, b)
}

func (b *PatchService) RegisterHTTP(ctx context.Context, router *mux.Router, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	if err := v1.RegisterPatchServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		panic(err)
	}
}

func (b *PatchService) CreatePatch(ctx context.Context, request *v1.CreatePatchRequest) (*v1.Patch, error) {

	b.temp = &v1.Patch{
		Id:          shortid.MustGenerate(),
		Name:        request.Name,
		Description: request.Description,
		Patch:       request.Patch,
		Permalink:   "",
	}

	return b.temp, nil
}

func (b *PatchService) GetPatch(ctx context.Context, request *v1.GetPatchRequest) (*v1.Patch, error) {
	return b.temp, nil
}
