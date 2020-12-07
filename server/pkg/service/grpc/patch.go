package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
	"github.com/warmans/thrl6p/server/pkg/filter"
	"github.com/warmans/thrl6p/server/pkg/metadata"
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
	fp, err := patch.Fingerprint(*p)
	if err != nil {
		return nil, ErrInternal(err).Err()
	}

	rec := &model.Patch{
		Name:        p.Data.Meta.Name,
		Fingerprint: fp,
		Description: req.Description,
		Patch:       req.Patch,
	}
	if err := b.db.WithStore(func(s *store.Store) error {
		exists, id, err := s.PatchExists(ctx, rec.Fingerprint)
		if err != nil {
			return err
		}
		if exists {
			return ErrInvalidRequestField("patch", "Identical patch already exists: "+id).Err()
		}
		return s.CreatePatch(rec)
	}); err != nil {
		return nil, ErrInternal(err).Err()
	}

	//todo: permalink
	return rec.Proto(""), nil
}

func (b *PatchService) GetPatch(ctx context.Context, request *v1.GetPatchRequest) (resp *v1.Patch, err error) {
	err = b.db.WithStore(func(s *store.Store) error {
		rec, err := s.GetPatch(ctx, request.Id)
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

func (b *PatchService) ListPatch(ctx context.Context, request *v1.ListPatchesRequest) (*v1.PatchList, error) {
	f, err := filter.Parse(request.Filter)
	if err != nil {
		return nil, ErrInvalidRequestField("filter", err.Error()).Err()
	}

	var patches []*model.Patch
	err = b.db.WithStore(func(s *store.Store) error {
		patches, err = s.ListPatches(ctx, f, request.PageSize, request.Page)
		return err
	})
	if err != nil {
		return nil, err
	}

	proto := &v1.PatchList{Patches: make([]*v1.Patch, len(patches))}
	for k, v := range patches {
		proto.Patches[k] = v.Proto("")
	}
	return proto, nil
}

func (b *PatchService) ValidateName(ctx context.Context, request *v1.ValidateNameRequest) (*v1.NameValidation, error) {
	var exists bool
	var id string
	err := b.db.WithStore(func(s *store.Store) error {
		var err error
		exists, id, err = s.PatchNameExists(ctx, request.Name)
		if err != nil {
			return ErrInternal(err).Err()
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if exists {
		return &v1.NameValidation{Ok: false, Reason: fmt.Sprintf("Name is not unique (duplicates %s)", id)}, nil
	}
	return &v1.NameValidation{Ok: true}, nil
}

func (b *PatchService) Metadata(ctx context.Context, _ *empty.Empty) (*v1.Meta, error) {
	return metadata.All().Proto(), nil
}
