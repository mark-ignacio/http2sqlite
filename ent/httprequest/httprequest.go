// Code generated by entc, DO NOT EDIT.

package httprequest

import (
	"time"
)

const (
	// Label holds the string label denoting the httprequest type in the database.
	Label = "http_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReceived holds the string denoting the received field in the database.
	FieldReceived = "received"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldHeader holds the string denoting the header field in the database.
	FieldHeader = "header"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// Table holds the table name of the httprequest in the database.
	Table = "http_requests"
)

// Columns holds all SQL columns for httprequest fields.
var Columns = []string{
	FieldID,
	FieldReceived,
	FieldHost,
	FieldPath,
	FieldMethod,
	FieldHeader,
	FieldBody,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultReceived holds the default value on creation for the "received" field.
	DefaultReceived func() time.Time
)
