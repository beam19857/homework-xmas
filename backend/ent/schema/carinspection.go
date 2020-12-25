package schema

import "github.com/facebookincubator/ent"

// Test
// CarInspection holds the schema definition for the CarInspection entity.
type CarInspection struct {
	ent.Schema
}

// Fields of the CarInspection.
func (CarInspection) Fields() []ent.Field {
	return nil
}

// Edges of the CarInspection.
func (CarInspection) Edges() []ent.Edge {
	return nil
}
