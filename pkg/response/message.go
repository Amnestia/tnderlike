package response

import "net/http"

// constant string for message/error
const (
	Success             = "Success"
	Created             = "Successfully created data"
	BadRequest          = "Bad request"
	Unathorized         = "Unathorized access"
	Forbidden           = "Forbidden"
	NotFound            = "Not found"
	UnprocessableEntity = "Unprocessable entity"
	InternalServerError = "Internal server error"
)

// ResponseMessage response message mapped to status code
var ResponseMessage = map[int]string{
	http.StatusOK:                  Success,
	http.StatusCreated:             Created,
	http.StatusBadRequest:          BadRequest,
	http.StatusUnauthorized:        Unathorized,
	http.StatusForbidden:           Forbidden,
	http.StatusNotFound:            NotFound,
	http.StatusUnprocessableEntity: UnprocessableEntity,
	http.StatusInternalServerError: InternalServerError,
}
