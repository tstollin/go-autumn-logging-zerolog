package loggermiddleware

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/rs/zerolog/log"
	"net/http"
)

var RequestIdFieldName = "request-id"
var MethodFieldName = "method"
var PathFieldName = "path"

// middleware to add context aware logger to your context
//
// also support logging request ids, just assign RequestIdFromContextRetriever before using this
func AddZerologLoggerToContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		method := r.Method
		path := r.URL.Path

		builder := log.Logger.With().
			Str(MethodFieldName, method).
			Str(PathFieldName, path)
		if aulogging.RequestIdRetriever != nil {
			requestId := aulogging.RequestIdRetriever(ctx)
			builder = builder.Str(RequestIdFieldName, requestId)
		}
		sublogger := builder.Logger()
		newCtx := sublogger.WithContext(ctx)

		next.ServeHTTP(w, r.WithContext(newCtx))
	}
	return http.HandlerFunc(fn)
}
