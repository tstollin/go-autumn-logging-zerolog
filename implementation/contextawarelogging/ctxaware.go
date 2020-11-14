package contextawarelogging

import (
	"github.com/StephanHCB/go-autumn-logging-zerolog/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
	"github.com/rs/zerolog"
)

type ZerologContextAwareLoggingImplementation struct{
	LoggerWithCtx *zerolog.Logger
}

func (ca *ZerologContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Trace()}
}

func (ca *ZerologContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Debug()}
}

func (ca *ZerologContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Info()}
}

func (ca *ZerologContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Warn()}
}

func (ca *ZerologContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Error()}
}

func (ca *ZerologContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Fatal()}
}

func (ca *ZerologContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{LeveledLogEvent: ca.LoggerWithCtx.Panic()}
}
