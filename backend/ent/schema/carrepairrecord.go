package schema

import "github.com/facebookincubator/ent"

// CarRepairRecord holds the schema definition for the CarRepairRecord entity.
type CarRepairRecord struct {
	ent.Schema
}

// Fields of the CarRepairRecord.
func (CarRepairRecord) Fields() []ent.Field {
	return nil
}

// Edges of the CarRepairRecord.
func (CarRepairRecord) Edges() []ent.Edge {
	return nil
}
