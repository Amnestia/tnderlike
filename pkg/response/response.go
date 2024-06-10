package response

import (
	"context"
	"net/http"

	"github.com/amnestia/tnderlike/pkg/json"
	"github.com/amnestia/tnderlike/pkg/logger"
)

// Response struct
type Response struct {
	Ctx        context.Context `json:"-"`
	Body       any             `json:"body,omitempty"`
	Error      string          `json:"error,omitempty"`
	Message    string          `json:"message"`
	StatusCode int             `json:"-"`
}

// NewResponse set new response with ctx
func NewResponse(ctx context.Context) *Response {
	return &Response{Ctx: ctx}
}

// SetMessage set Response.Message
func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

// SetBody set Response.Body
func (r *Response) SetBody(body any) *Response {
	r.Body = body
	return r
}

// SetError set Response.Error
func (r *Response) SetError(err string) *Response {
	r.Error = err
	return r
}

// SetStatusCode set Response.StatusCode
func (r *Response) SetStatusCode(code int) *Response {
	r.StatusCode = code
	return r
}

// SetResponse set response attribute
func (r *Response) SetResponse(code int, body any, msg string) *Response {
	return r.SetBody(body).SetMessage(msg).SetStatusCode(code)
}

func (r *Response) WriteJSON(w http.ResponseWriter) {
	b, _ := json.Marshal(r.Return())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	_, err := w.Write(b)
	if err != nil {
		logger.Logger.Error().Err(logger.ErrorWrap(err, "WriteJSON.Write"))
	}
}

// SetErrorResponse set error response attribute
func (r *Response) SetErrorResponse(code int, err error, msg ...string) *Response {
	message := err.Error()
	if len(msg) > 0 {
		message = msg[0]
	}
	logger.Logger.Error().Err(logger.ErrorWrap(err, ""))
	return r.SetMessage(message).SetStatusCode(code)
}

// Return add default error and message for response
func (r *Response) Return() *Response {
	if r.Message == "" {
		r.SetMessage(ResponseMessage[r.StatusCode])
	}
	if r.StatusCode/100 != 2 {
		if r.Error == "" {
			r.SetError(ResponseMessage[r.StatusCode])
		}
	}
	if r.StatusCode/100 == 5 {
		r.SetMessage("Internal Server Error")
	}
	return r
}
