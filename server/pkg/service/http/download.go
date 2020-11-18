package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

type DownloadService struct {

}

func (c *DownloadService) RegisterHTTP(ctx context.Context, router *mux.Router) {
	router.Path("/dl/{id}").Handler(http.HandlerFunc(c.PatchDownloadHandler))
}

func (c *DownloadService) PatchDownloadHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("A PATCH"))
}
