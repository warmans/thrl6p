package grpc

import (
	"context"
	"encoding/json"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
	"github.com/warmans/thrl6p/server/pkg/patch"
	"github.com/warmans/thrl6p/server/pkg/store"
	"github.com/warmans/thrl6p/server/pkg/store/model"
	"google.golang.org/grpc"
)

func NewPatchService(db *store.Conn) *PatchService {
	return &PatchService{db: db}
}

type PatchService struct {
	db *store.Conn
}

func (b *PatchService) RegisterGRPC(server *grpc.Server) {
	v1.RegisterPatchServiceServer(server, b)
}

func (b *PatchService) RegisterHTTP(ctx context.Context, router *mux.Router, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	if err := v1.RegisterPatchServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		panic(err)
	}
}

func (b *PatchService) CreatePatch(ctx context.Context, req *v1.CreatePatchRequest) (*v1.Patch, error) {

	p := &patch.THRL6P{}
	if err := json.Unmarshal([]byte(req.Patch), p); err != nil {
		return nil, ErrInvalidRequestField("patch", err.Error()).Err()
	}

	rec := &model.Patch{
		Name:        p.Data.Meta.Name,
		Description: req.Patch,
		Patch:       req.Patch,
	}
	if err := b.db.WithStore(func(s *store.Store) error {
		return s.CreatePatch(rec)
	}); err != nil {
		return nil, ErrInternal().Err()
	}

	//todo: permalink
	return rec.Proto(""), nil
}

func (b *PatchService) GetPatch(ctx context.Context, request *v1.GetPatchRequest) (resp *v1.Patch, err error) {
	err = b.db.WithStore(func(s *store.Store) error {
		rec, err := s.GetPatch(request.Id)
		if err != nil {
			return err
		}
		resp = rec.Proto("") //todo: permalink
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (b *PatchService) ValidateName(ctx context.Context, request *v1.ValidateNameRequest) (*empty.Empty, error) {
	return nil, ErrInvalidRequestField("name", "Name is not unique").Err()
}
