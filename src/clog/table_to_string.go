package clog

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/piot/log-go/src/clogint"
)

func getColumnWidth(r []string, widths []int) {
	for columnIndex, c := range r {
		current := widths[columnIndex]
		widthWithSpace := len(c)
		if widthWithSpace > current {
			widths[columnIndex] = widthWithSpace
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

func trimString(s string, maxWidth int) string {
	if len(s) > maxWidth {
		return s[:maxWidth-3] + "..."
	}
	return s
}

func expandTableStringers(stringers [][]fmt.Stringer) [][]string {
	targetRows := make([][]string, len(stringers))

	for rowIndex, r := range stringers {
		var targetColumns []string

		for _, c := range r {
			targetColumns = append(targetColumns, c.String())
		}

		targetRows[rowIndex] = targetColumns
	}

	return targetRows
}

func tableToString(t clogint.Table, colorIndex int) string {
	dataToUse := t.Data
	if t.DataStringer != nil {
		dataToUse = expandTableStringers(t.DataStringer)
	}
	columnWidths := make([]int, len(t.Columns))

	getColumnWidth(t.Columns, columnWidths)
	err := getColumnWidths(dataToUse, columnWidths)
	if err != nil {
		return err.Error()
	}

	//	columnColor := color.New(color.FgHiYellow).Add(color.BgHiBlue)
	colorPalette := []*color.Color{
		color.New(color.FgHiYellow).Add(color.BgBlue),
		color.New(color.FgHiWhite).Add(color.BgBlue),

		color.New(color.FgHiYellow).Add(color.BgMagenta),
		color.New(color.FgHiWhite).Add(color.BgMagenta),

		color.New(color.FgHiYellow).Add(color.BgCyan),
		color.New(color.FgHiWhite).Add(color.BgCyan),
	}

	columnPalette := []*color.Color{
		color.New(color.FgBlack).Add(color.BgHiBlue),
		color.New(color.FgBlack).Add(color.BgHiBlue),

		color.New(color.FgBlack).Add(color.BgHiMagenta),
		color.New(color.FgBlack).Add(color.BgHiMagenta),

		color.New(color.FgBlack).Add(color.BgHiCyan),
		color.New(color.FgBlack).Add(color.BgHiCyan),
	}

	dataColor := colorPalette[colorIndex%len(colorPalette)]
	columnColor := columnPalette[colorIndex%len(columnPalette)]

	sum := 0
	for _, width := range columnWidths {
		sum += width
	}

	const clogIndentAndPadding = 4

	const maxScreenWidth = 132 - clogIndentAndPadding

	if sum > maxScreenWidth {
		overflow := sum - maxScreenWidth
		removeFromEachColumn := overflow / len(columnWidths)

		for i := range columnWidths {
			if columnWidths[i] > removeFromEachColumn+2 {
				columnWidths[i] -= removeFromEachColumn
			}
		}
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
	columnS += "┌ "

	for index, col := range columnNames {
		columnString := fmt.Sprintf(columnFormat[index], trimString(col, columnWidths[index]))

		if index > 0 {
			columnS += " ┬ "
		}

		columnS += columnString
	}

	columnS += " ┐"
	dataS := ""

	for rowIndex, row := range dataToUse {
		if rowIndex > 0 {
			dataS += "\n"
		}
		dataRowS := "| "

		for index, col := range row {
			columnString := fmt.Sprintf(columnFormat[index], trimString(col, columnWidths[index]))

			if index > 0 {
				dataRowS += " | "
			}

			dataRowS += columnString
		}

		dataRowS += " |"

		dataS += dataColor.Sprint(dataRowS)
	}
	s := "\n" + columnColor.Sprint(columnS) + "\n" + dataS

	return s
}
