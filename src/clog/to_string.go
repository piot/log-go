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
		return fmt.Sprintf("%v", f.Integer)
	case clogint.StringType:
		return fmt.Sprintf("'%v'", f.String)
	case clogint.StringerType:
		return fmt.Sprintf("'%v'", f.Other)
	case clogint.InterfaceType:
		return fmt.Sprintf("'%v'", f.Other)
	case clogint.BoolType:
		return fmt.Sprintf("%v", f.Integer)
	case clogint.BlobType:
		hexString := hex.Dump(f.Other.([]byte))
		return hexString
	case clogint.ErrorType:
		return fmt.Sprintf("'%s'", f.Other.(error).Error())
	}
	panic("unknown type")
}
