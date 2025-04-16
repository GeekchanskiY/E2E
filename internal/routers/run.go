package routers

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Run(r *Router) error {
	r.setup()

	r.setupFileServer()

	r.addRoutes()

	go func() {
		r.logger.Info(
			"running server",
			zap.String("addr", fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)),
		)

		server := &http.Server{
			Addr:              fmt.Sprintf("%s:%d", r.config.Host, r.config.Port),
			ReadHeaderTimeout: 3 * time.Second,
			Handler:           r.mux,
		}

		err := server.ListenAndServe()
		if err != nil {
			r.logger.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	return nil
}
