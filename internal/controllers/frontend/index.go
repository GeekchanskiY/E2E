package frontend

import "net/http"

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("frontend.index")

}
