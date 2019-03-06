package clogint

type FieldType uint8

const (
	IntType FieldType = iota
	UintType
	StringType
	BoolType
	ErrorType
)

type LogLevel uint8

const (
	Trace LogLevel = iota
	Debug
	Info
	Warning
	Error
)

type Field struct {
	Type    FieldType
	Key     string
	Integer int64
	String  string
	Other   interface{}
}
