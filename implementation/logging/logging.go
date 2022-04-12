package logging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-zerolog/implementation/contextawarelogging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"github.com/rs/zerolog/log"
)

type ZerologLoggingImplementation struct{}

func (l *ZerologLoggingImplementation) Ctx(ctx context.Context) auloggingapi.ContextAwareLoggingImplementation {
	return &contextawarelogging.ZerologContextAwareLoggingImplementation{
		LoggerWithCtx: log.Ctx(ctx),
		Ctx:           ctx,
	}
}

func (l *ZerologLoggingImplementation) NoCtx() auloggingapi.ContextAwareLoggingImplementation {
	return &contextawarelogging.ZerologContextAwareLoggingImplementation{
		LoggerWithCtx: &log.Logger,
		Ctx:           context.Background(),
	}
}
