package schema

import "github.com/facebookincubator/ent"

// Ambulance holds the schema definition for the Ambulance entity.
type Ambulance struct {
	ent.Schema
}
//this is homework
// Fields of the Ambulance.
func (Ambulance) Fields() []ent.Field {
	return nil
}

// Edges of the Ambulance.
func (Ambulance) Edges() []ent.Edge {
	return nil
}
