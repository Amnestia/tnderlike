package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/amnestia/tnderlike/pkg/logger"
	"github.com/amnestia/tnderlike/pkg/response"
	"github.com/go-chi/chi/v5/middleware"
)

// Logger log api request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		next.ServeHTTP(ww, r)
		time.Since(start)
	})
}

// PanicRecovery panic recovery handler
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error().Err(logger.ErrorWrap(fmt.Errorf("panic occurred on %v: %v : %v\r\n", r.URL.Path, err, string(debug.Stack())), ""))
				response.NewResponse(r.Context()).SetResponse(http.StatusInternalServerError, nil, "").WriteJSON(w)
				if err != nil {
					logger.Logger.Error().Err(logger.ErrorWrap(fmt.Errorf("error on recovering %v : %v", r.URL.Path, err), "PanicRecovery.recover"))
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
