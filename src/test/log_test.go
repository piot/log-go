package clogtest

import (
	"fmt"
	"testing"

	"github.com/piot/log-go/src/clog"
)

func Test(t *testing.T) {
	log := clog.DefaultLog()

	log.Debug("Just some debugging")
	log.Info("Boot up modem", clog.String("name", "clarios"))
	log.Warn("I have to tell you something", clog.String("greeting", "something"), clog.Int("code", -42))
	log.Trace("This happens a lot", clog.Bool("engine_is_on", true))
	log.Err(fmt.Errorf("some generic error message"))
}
