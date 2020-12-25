package schema

import "github.com/facebookincubator/ent"

// CarRegister holds the schema definition for the CarRegister entity.
type CarRegister struct {
	ent.Schema
}

// Fields of the CarRegister.
func (CarRegister) Fields() []ent.Field {
	return nil
}

// Edges of the CarRegister.
func (CarRegister) Edges() []ent.Edge {
	return nil
}
