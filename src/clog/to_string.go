package clog

import (
	"encoding/hex"
	"fmt"

	"github.com/piot/log-go/src/clogint"
)

func ConvertToString(fields []Field) string {
	s := "{"

	for index, f := range fields {
		if index > 0 {
			s += ", "
		}
		s += f.Key + "=" + fieldValueToString(f)
	}

	s += "}"

	return s
}

func fieldValueToString(f Field) string {
	switch f.Type {
	case clogint.IntType:
		return fmt.Sprintf("%v", f.Integer)
	case clogint.UintType:
		return fmt.Sprintf("%v", f.UnsignedInteger)
	case clogint.StringType:
		return fmt.Sprintf("'%v'", f.String)
	case clogint.StringSliceType:
		s := ""

		for index, stringValue := range f.Other.([]string) {
			if index > 0 {
				s += ", "
			}
			s += fmt.Sprintf("'%v'", stringValue)
		}

		return s
	case clogint.StringerType:
		return fmt.Sprintf("'%v'", f.Other)
	case clogint.StringerSliceType:
		s := ""

		for index, stringer := range f.Other.([]fmt.Stringer) {
			if index > 0 {
				s += ", "
			}
			s += fmt.Sprintf("'%v'", stringer)
		}

		return s
	case clogint.InterfaceType:
		return fmt.Sprintf("'%v'", f.Other)
	case clogint.BoolType:
		return fmt.Sprintf("%v", f.Integer)
	case clogint.BlobType:
		hexString := hex.Dump(f.Other.([]byte))
		return hexString
	case clogint.TableType:
		{
			t := f.Other.(clogint.Table)

			return tableToString(t, f.ColorIndex)
		}
	case clogint.ErrorType:
		return fmt.Sprintf("'%s'", f.Other.(error).Error())
	}

	panic("log-go: unknown type")
}
