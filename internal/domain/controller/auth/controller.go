package auth

import (
	"fmt"
	"net/http"

	authmodel "github.com/amnestia/tnderlike/internal/domain/model/auth"
	"github.com/amnestia/tnderlike/internal/domain/service"
	"github.com/amnestia/tnderlike/pkg/json"
	"github.com/amnestia/tnderlike/pkg/response"
)

// Controller handler for this domain
type Controller struct {
	AuthSvc service.AuthServicer
}

// Auth authentication login handler
func (c *Controller) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &authmodel.Account{}
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err, "Invalid Request").WriteJSON(w)
		return
	}
	if req.Email == "" {
		response.SetErrorResponse(http.StatusBadRequest, fmt.Errorf("Email is required")).WriteJSON(w)
		return
	}
	if req.Password == "" {
		response.SetErrorResponse(http.StatusNotFound, fmt.Errorf("Password is required")).WriteJSON(w)
		return
	}
	ret := c.AuthSvc.Auth(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
	response.SetResponse(ret.HTTPCode, ret.TokenData, "Successfully logged in").WriteJSON(w)
}

// Register register new account
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &authmodel.Account{}
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err).WriteJSON(w)
		return
	}
	err = validateRegister(req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err).WriteJSON(w)
		return
	}
	ret := c.AuthSvc.Register(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
	response.SetResponse(ret.HTTPCode, nil, "successfully registered").WriteJSON(w)
}
