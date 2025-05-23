package routers

import (
	"net/http"

	"finworker/static"
)

func (r *Router) setupFileServer() {
	r.mux.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(static.Fs))))
	r.mux.Handle("/media/*", http.StripPrefix("/media", http.FileServer(http.Dir("media"))))
}
