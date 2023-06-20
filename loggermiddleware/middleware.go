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

type CustomJsonLogField struct {
	LogFieldName   string
	ValueExtractor func(r *http.Request) string
}

type AddZerologLoggerToContextOptions struct {
	CustomJsonLogFields []CustomJsonLogField
}

// AddZerologLoggerToContextMiddleware constructs a middleware with the given options.
//
// If you don't need to set custom options, you can just directly use AddZerologLoggerToContext as your middleware
// instead.
func AddZerologLoggerToContextMiddleware(options AddZerologLoggerToContextOptions) func(http.Handler) http.Handler {
	mw := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			method := r.Method
			path := r.URL.Path

			builder := log.Logger.With()
			if auzerolog.IsJson {
				builder = builder.
					Str(MethodFieldName, method).
					Str(PathFieldName, path)

				for _, customField := range options.CustomJsonLogFields {
					value := customField.ValueExtractor(r)
					builder = builder.Str(customField.LogFieldName, value)
				}
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
	return mw
}

var defaultOptionsInstance = AddZerologLoggerToContextMiddleware(AddZerologLoggerToContextOptions{})

// AddZerologLoggerToContext is a middleware to add a context aware logger to your context.
//
// also supports logging request ids, just assign RequestIdFromContextRetriever before using this
func AddZerologLoggerToContext(next http.Handler) http.Handler {
	return defaultOptionsInstance(next)
}

// AddCustomJsonLogField is a function to add custom fields to json logs.
//
// DEPRECATED. This function is deprecated and calling it no longer has any effect!
//
// Please use AddZerologLoggerToContextMiddleware and an options object to pass options.
func AddCustomJsonLogField(logFieldName string, valueExtractor func(r *http.Request) string) {
	// functionality intentionally removed
}
