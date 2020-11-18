package util

import (
	"net/http"
	"os"
	"path"
)

func TryFileHandler(defaultFile string, fileDirs ...string) http.Handler {
	return &TryFiles{fileDirs: fileDirs, defaultFile: defaultFile}
}

type TryFiles struct {
	fileDirs    []string
	defaultFile string
}

func (h *TryFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "" {
		for _, v := range h.fileDirs {
			filePath := path.Join(v, r.URL.Path)
			_, err := os.Stat(filePath)
			if err == nil {
				http.ServeFile(rw, r, filePath)
				return
			}
		}
	}
	http.ServeFile(rw, r, h.defaultFile)
}