package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").Default("unknown"),
	}
}

func (User) Annotations() []schema.Annotation {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(
			entgql.Skip(
				entgql.SkipMutationCreateInput,
				entgql.SkipMutationUpdateInput,
			)),
	}
}

func (User) Indexes() []ent.Index {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
