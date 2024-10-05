package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DefaulTimeMixin struct {
	mixin.Schema
}

func (DefaulTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").
			Immutable().
			Default(time.Now),

		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
