package routers

import (
	"net/http"

	"finworker/internal/static"
)

func (r *Router) setupFileServer() {
	r.mux.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(static.Fs))))
}
