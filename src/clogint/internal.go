package clogint

import "fmt"

type FieldType uint8

const (
	IntType FieldType = iota
	UintType
	StringType
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
	Other           interface{}
}

type Table struct {
	Columns      []string
	Data         [][]string
	DataStringer [][]fmt.Stringer
}
