package auzerolog

import (
	"bytes"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var RequestIdFieldName = "request-id"

// SetupJsonLogging configures go-autumn-logging to log via rs/zerolog using a subset of ECS
// see https://www.elastic.co/guide/en/ecs/1.4
func SetupJsonLogging(serviceId string) {
	zerolog.TimestampFieldName = "@timestamp"
	zerolog.LevelFieldName = "log.level"
	zerolog.MessageFieldName = "message" // correct by default

	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service.id", serviceId).
		Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	IsJson = true
}

var IsJson = false

func plaintextFieldNameOmitter(name interface{}) string {
	return ""
}

func plaintextFieldValueOmitter(value interface{}) string {
	return ""
}

func plaintextRequestIdFormatter(value interface{}) string {
	vStr, ok := value.(string)
	if !ok {
		return aulogging.DefaultRequestIdValue
	}
	return "[" + vStr + "]"
}

// SetupPlaintextLogging configures go-autumn-logging to log via rs/zerolog using a simple plaintext logger
func SetupPlaintextLogging() {
	zerolog.CallerFieldName = RequestIdFieldName // avoid double printing the request id (bit of a hack)

	log.Logger = log.With().Str(RequestIdFieldName, aulogging.DefaultRequestIdValue).Logger().Output(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    true,
			TimeFormat: "15:04:05.000",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				RequestIdFieldName,
				zerolog.MessageFieldName,
			},
			FormatFieldName:  plaintextFieldNameOmitter,
			FormatFieldValue: plaintextFieldValueOmitter,
			FormatCaller:     plaintextRequestIdFormatter,
		})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	IsJson = false
}

var RecordedLogForTesting = new(bytes.Buffer)

// Setup function for testing that records log entries instead of writing them to console
func SetupForTesting() {
	SetupPlaintextLogging()
	log.Logger = zerolog.New(RecordedLogForTesting).With().Timestamp().Logger()
}
