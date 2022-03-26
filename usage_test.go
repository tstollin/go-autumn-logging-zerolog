package go_autumn_logging_zerolog

import (
	"context"
	"errors"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"testing"
)

func TestPlaintextLogging(t *testing.T) {
	SetupPlaintextLogging()

	ctx := AddLoggerToCtx(context.Background())
	err := errors.New("some wonderful error")
	
	aulogging.Logger.NoCtx().Error().Print("no context print error severity, no error object")
	aulogging.Logger.Ctx(ctx).Error().WithErr(err).Print("with context print error severity, with error object")

	aulogging.Logger.NoCtx().Warn().Print("no context print warn severity")
	aulogging.Logger.Ctx(ctx).Warn().Print("with context print warn severity")

	aulogging.Logger.NoCtx().Info().Print("no context print info severity")
	aulogging.Logger.Ctx(ctx).Info().Print("with context print info severity")

	aulogging.Logger.NoCtx().Debug().Print("no context print debug severity (should not show)")
	aulogging.Logger.Ctx(ctx).Debug().Print("with context print debug severity (should not show)")

	aulogging.Logger.NoCtx().Trace().Print("no context print trace severity (should not show)")
	aulogging.Logger.Ctx(ctx).Trace().Print("with context print trace severity (should not show)")
}

func TestJsonLogging(t *testing.T) {
	SetupJsonLogging("my-service")

	ctx := AddLoggerToCtx(context.Background())
	err := errors.New("some wonderful error")

	aulogging.Logger.NoCtx().Error().Print("no context print error severity, no error object")
	aulogging.Logger.Ctx(ctx).Error().WithErr(err).Print("with context print error severity, with error object")

	aulogging.Logger.NoCtx().Warn().Print("no context print warn severity")
	aulogging.Logger.Ctx(ctx).Warn().Print("with context print warn severity")

	aulogging.Logger.NoCtx().Info().Print("no context print info severity")
	aulogging.Logger.Ctx(ctx).Info().Print("with context print info severity")

	aulogging.Logger.NoCtx().Debug().Print("no context print debug severity (should not show)")
	aulogging.Logger.Ctx(ctx).Debug().Print("with context print debug severity (should not show)")

	aulogging.Logger.NoCtx().Trace().Print("no context print trace severity (should not show)")
	aulogging.Logger.Ctx(ctx).Trace().Print("with context print trace severity (should not show)")
}
