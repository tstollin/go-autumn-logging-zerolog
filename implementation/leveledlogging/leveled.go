package leveledlogging

import (
	"context"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"github.com/rs/zerolog"
)

type ZerologLeveledLoggingImplementation struct {
	LeveledLogEvent *zerolog.Event
	Ctx             context.Context
	Level           string
	Err             error
	Additional      map[string]string
}

func (lv *ZerologLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	lv.LeveledLogEvent = lv.LeveledLogEvent.Err(err)
	lv.Err = err
	return lv
}

func (lv *ZerologLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	lv.LeveledLogEvent = lv.LeveledLogEvent.Str(key, value)
	if lv.Additional == nil {
		lv.Additional = make(map[string]string, 0)
	}
	lv.Additional[key] = value
	return lv
}

func (lv *ZerologLeveledLoggingImplementation) Print(v ...interface{}) {
	message := fmt.Sprint(v...)
	if aulogging.LogEventCallback != nil {
		aulogging.LogEventCallback(lv.Ctx, lv.Level, message, lv.Err, lv.Additional)
	}
	lv.LeveledLogEvent.Msg(message)
}

func (lv *ZerologLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if aulogging.LogEventCallback != nil {
		aulogging.LogEventCallback(lv.Ctx, lv.Level, message, lv.Err, lv.Additional)
	}
	lv.LeveledLogEvent.Msg(message)
}
