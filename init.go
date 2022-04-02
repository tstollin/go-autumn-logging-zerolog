package auzerolog

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-zerolog/implementation/logging"
)

func init() {
	aulogging.Logger = &logging.ZerologLoggingImplementation{}
}
