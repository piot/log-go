package clogtest

import (
	"fmt"
	"testing"

	"github.com/piot/log-go/src/clog"
)

type TestStruct struct {
}

func (t TestStruct) String() string {
	return "stringed!"
}
func Test(t *testing.T) {
	log := clog.DefaultLog()

	log.Debug("Just some debugging")
	log.Info("Boot up modem", clog.String("name", "clarios"))
	log.Warn("I have to tell you something", clog.String("greeting", "something"), clog.Int("code", -42))
	log.Trace("This happens a lot", clog.Bool("engine_is_on", true))
	log.Trace("stringer", clog.Stringer("test", TestStruct{}))
	log.Err(fmt.Errorf("some generic error message"))
	log.Error("stringer", clog.Stringer("error", TestStruct{}))
}

func TestFile(t *testing.T) {
	log, logErr := clog.DefaultFileLog("CLog", "testing")
	if logErr != nil {
		t.Fatal(logErr)
	}

	log.Debug("Just some debugging")
	log.Info("Boot up modem", clog.String("name", "clarios"))
	log.Warn("I have to tell you something", clog.String("greeting", "something"), clog.Int("code", -42))
	log.Trace("This happens a lot", clog.Bool("engine_is_on", true))
	log.Trace("stringer", clog.Stringer("test", TestStruct{}))
	log.Err(fmt.Errorf("some generic error message"))
	log.Error("stringer", clog.Stringer("error", TestStruct{}))
}
