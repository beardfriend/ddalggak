package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entsql.WithComments(true),
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),

		field.String("password").Sensitive(),

		field.String("nickname"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaulTimeMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}
