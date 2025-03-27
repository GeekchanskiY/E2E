package routers

import (
	"fmt"
	"net/http"

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

		err := http.ListenAndServe(fmt.Sprintf("%s:%d", r.config.Host, r.config.Port), r.mux)
		if err != nil {
			r.logger.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	return nil
}
