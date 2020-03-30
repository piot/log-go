package clog

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
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

func getColumnWidth(r []string, widths []int) {
	for columnIndex, c := range r {
		current := widths[columnIndex]
		if len(c) > current {
			widths[columnIndex] = len(c)
		}
	}
}

func getColumnWidths(data [][]string, columnWidths []int) error {
	for rowIndex, r := range data {
		if len(r) != len(columnWidths) {
			return fmt.Errorf("row index %v has diffent column count:%v, expected %v", rowIndex,
				len(r), len(columnWidths))
		}
		getColumnWidth(r, columnWidths)
	}

	return nil
}

func tableToString(t clogint.Table) string {
	columnWidths := make([]int, len(t.Columns))

	getColumnWidth(t.Columns, columnWidths)
	err := getColumnWidths(t.Data, columnWidths)
	if err != nil {
		return err.Error()
	}

	columnFormat := make([]string, len(t.Columns))
	columnNames := make([]string, len(t.Columns))
	for index, col := range t.Columns {
		width := columnWidths[index]

		prefix := "-"
		columnNames[index] = col
		if strings.HasPrefix(col, "-") || strings.HasPrefix(col, "+") {
			prefix = col[0:1]
			columnNames[index] = col[1:]
		}

		fmtString := "%" + prefix + strconv.Itoa(width) + "s"
		columnFormat[index] = fmtString
	}

	columnS := ""
	columnS += "\n|"
	for index, col := range columnNames {
		columnString := fmt.Sprintf(columnFormat[index], col)
		if index > 0 {
			columnS += "|"
		}
		columnS += columnString
	}
	columnS += "|\n"

	dataS := ""

	for _, row := range t.Data {
		dataS += "|"

		for index, col := range row {
			columnString := fmt.Sprintf(columnFormat[index], col)
			if index > 0 {
				dataS += "|"
			}
			dataS += columnString
		}

		dataS += "|"
	}

	s := color.CyanString(columnS) + color.HiMagentaString(dataS)

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
		return tableToString(f.Other.(clogint.Table))
	case clogint.ErrorType:
		return fmt.Sprintf("'%s'", f.Other.(error).Error())
	}
	panic("unknown type")
}
