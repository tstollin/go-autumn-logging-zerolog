package go_autumn_logging_zerolog

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// configure to implement a subset of ECS
// see https://www.elastic.co/guide/en/ecs/1.4
func SetupJsonLogging(serviceId string) {
	zerolog.TimestampFieldName = "@timestamp"
	zerolog.LevelFieldName = "log.level"
	zerolog.MessageFieldName = "message" // correct by default

	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service.id", serviceId).
		Logger()
}

// configure to implement a simple plaintext logger
func SetupPlaintextLogging() {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stdout,
			NoColor: true,
			TimeFormat: "15:04:05.000",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.MessageFieldName,
			},
		})
}

var RecordedLogForTesting = new(bytes.Buffer)

// Setup function for testing that records log entries instead of writing them to console
func SetupForTesting() {
	SetupPlaintextLogging()
	log.Logger = zerolog.New(RecordedLogForTesting).With().Timestamp().Logger()
}
