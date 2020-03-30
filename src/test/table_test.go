package clogtest

import (
	"testing"

	"github.com/piot/log-go/src/clog"
)

func TestTable(t *testing.T) {
	log := clog.DefaultLog()

	log.Info("hey", clog.Table("data", []string{"+firstHeader", "lastColumn"}, [][]string{{"abcdefgh", "This is just a long string"}}))
	log.Warn("completely different")
}
