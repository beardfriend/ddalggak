package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "products"},
		entsql.WithComments(true),
	}
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),

		field.Int("userID"),

		field.String("price"),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaulTimeMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("products").
			Unique().
			Required().
			Field("userID"),
	}
}
