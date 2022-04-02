package loggermiddleware

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	auzerolog "github.com/StephanHCB/go-autumn-logging-zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

var RequestIdFieldName = "request-id"
var MethodFieldName = "method"
var PathFieldName = "path"

// AddZerologLoggerToContext is a middleware to add a context aware logger to your context.
//
// also supports logging request ids, just assign RequestIdFromContextRetriever before using this
func AddZerologLoggerToContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		method := r.Method
		path := r.URL.Path

		builder := log.Logger.With()
		if auzerolog.IsJson {
			builder = builder.
				Str(MethodFieldName, method).
				Str(PathFieldName, path)
		}
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
