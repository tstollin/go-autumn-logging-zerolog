package leveledlogging

import (
	"fmt"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"github.com/rs/zerolog"
)

type ZerologLeveledLoggingImplementation struct {
	LeveledLogEvent  *zerolog.Event
}

func (lv *ZerologLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	lv.LeveledLogEvent = lv.LeveledLogEvent.Err(err)
	return lv
}

func (lv *ZerologLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	lv.LeveledLogEvent = lv.LeveledLogEvent.Str(key, value)
	return lv
}

func (lv *ZerologLeveledLoggingImplementation) Print(v ...interface{}) {
	lv.LeveledLogEvent.Msg(fmt.Sprint(v...))
}

func (lv *ZerologLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
	lv.LeveledLogEvent.Msgf(format, v...)
}
