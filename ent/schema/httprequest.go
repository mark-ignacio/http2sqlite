package schema

import (
	"net/http"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HTTPRequest holds the schema definition for the HTTPRequest entity.
type HTTPRequest struct {
	ent.Schema
}

// Fields of the HTTPRequest.
func (HTTPRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Time("received").Default(time.Now),
		field.String("host"),
		field.String("path"),
		field.String("method"),
		field.JSON("header", http.Header{}),
		field.Bytes("body"),
	}
}

// Edges of the HTTPRequest.
func (HTTPRequest) Edges() []ent.Edge {
	return nil
}
