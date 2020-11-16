package grpc

import (
	"context"
	"encoding/json"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/teris-io/shortid"
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
	"github.com/warmans/thrl6p/server/pkg/patch"
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

	p := &patch.THRL6P{}
	if err := json.Unmarshal([]byte(request.Patch), p); err != nil {
		return nil, ErrInvalidRequestField("patch", err.Error()).Err()
	}

	b.temp = &v1.Patch{
		Id:          shortid.MustGenerate(),
		Name:        p.Data.Meta.Name,
		Description: request.Description,
		Patch:       request.Patch,
		Permalink:   "",
	}

	return b.temp, nil
}

func (b *PatchService) GetPatch(ctx context.Context, request *v1.GetPatchRequest) (*v1.Patch, error) {
	return b.temp, nil
}

func (b *PatchService) ValidateName(ctx context.Context, request *v1.ValidateNameRequest) (*empty.Empty, error) {
	if b.temp == nil {
		return &empty.Empty{}, nil
	}
	if request.Name == b.temp.Name {
		return nil, ErrInvalidRequestField("name", "Name is not unique").Err()
	}
	return &empty.Empty{}, nil
}
