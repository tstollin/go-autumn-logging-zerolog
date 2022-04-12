package contextawarelogging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-zerolog/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
	"github.com/rs/zerolog"
)

type ZerologContextAwareLoggingImplementation struct {
	LoggerWithCtx *zerolog.Logger
	Ctx           context.Context
}

func (ca *ZerologContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Trace(),
		Ctx:             ca.Ctx,
		Level:           "TRACE",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Debug(),
		Ctx:             ca.Ctx,
		Level:           "DEBUG",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Info(),
		Ctx:             ca.Ctx,
		Level:           "INFO",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Warn(),
		Ctx:             ca.Ctx,
		Level:           "WARN",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Error(),
		Ctx:             ca.Ctx,
		Level:           "ERROR",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Fatal(),
		Ctx:             ca.Ctx,
		Level:           "FATAL",
	}
}

func (ca *ZerologContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.ZerologLeveledLoggingImplementation{
		LeveledLogEvent: ca.LoggerWithCtx.Panic(),
		Ctx:             ca.Ctx,
		Level:           "PANIC",
	}
}
