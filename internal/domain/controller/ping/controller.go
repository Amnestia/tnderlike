package ping

import (
	"net/http"

	"github.com/amnestia/tnderlike/internal/domain/service"
	"github.com/amnestia/tnderlike/pkg/response"
)

// Controller handler for this domain
type Controller struct {
	PingSvc service.PingServicer
}

// Ping healthy check ping handler
func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ret := c.PingSvc.Ping()
	response.NewResponse(ctx).SetResponse(http.StatusOK, ret, "Success").WriteJSON(w)
}
