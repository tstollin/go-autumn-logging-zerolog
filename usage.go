package auzerolog

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

func AddLoggerToCtx(ctx context.Context) context.Context {
	return log.Logger.WithContext(ctx)
}
