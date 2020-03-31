package clogtest

import (
	"fmt"
	"testing"

	"github.com/piot/log-go/src/clog"
)

type SomeStringer struct {
}

func (s SomeStringer) String() string {
	return fmt.Sprintf("Resolved!")
}

func TestTable(t *testing.T) {
	log := clog.DefaultLog()

	log.Info("table log", clog.Table("a table", []string{"+firstColumn", "lastColumn"}, [][]string{{"First test", "This is just a long string"}}))
	log.Info("table log", clog.TableStringer("a table", []string{"+firstColumn", "lastColumn"}, [][]fmt.Stringer{{SomeStringer{}, SomeStringer{}}}))
	log.Warn("completely different")
}
