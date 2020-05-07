package clogint

import "fmt"

type FieldType uint8

const (
	IntType FieldType = iota
	UintType
	StringType
	StringSliceType
	StringerType
	StringerSliceType
	InterfaceType
	BoolType
	BlobType
	ErrorType
	TableType
)

type LogLevel uint8

const (
	Trace LogLevel = iota
	Debug
	Info
	Warning
	Error
	Panic
)

type Field struct {
	Type            FieldType
	Key             string
	Integer         int64
	UnsignedInteger uint64
	String          string
	ColorIndex      int
	Other           interface{}
}

func (f Field) WithColorIndex(index int) Field {
	newF := f
	newF.ColorIndex = index

	return newF
}

type Table struct {
	Columns      []string
	Data         [][]string
	DataStringer [][]fmt.Stringer
}
