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

var customJsonLogFields = make([]customJsonLogField, 0)

type customJsonLogField struct {
	LogFieldName   string
	ValueExtractor func(r *http.Request) string
}

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

			for _, customField := range customJsonLogFields {
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

// AddCustomJsonLogField is a function to add custom fields to json logs.
//
// The first parameter determines the name of the field. The second parameters is used to determine the value of the
// field from the logged request.
// Each invocation adds one extra field.
func AddCustomJsonLogField(logFieldName string, valueExtractor func(r *http.Request) string) {
	customJsonLogFields = append(customJsonLogFields, customJsonLogField{
		LogFieldName:   logFieldName,
		ValueExtractor: valueExtractor,
	})
}
