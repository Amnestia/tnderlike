package common

// DefaultResponse default response value that should be return from usecase
type DefaultResponse struct {
	HTTPCode int
	Error    error
}

// ListResponse response containing current page and total page
type ListResponse struct {
	Data        any   `json:"data"`
	CurrentPage int64 `json:"current_page"`
	TotalPage   int64 `json:"total_page"`
	DefaultResponse
}

// Build helper for building new response
func (d *DefaultResponse) Build(httpCode int, err error) *DefaultResponse {
	d.HTTPCode = httpCode
	d.Error = err
	return d
}
