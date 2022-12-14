package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		foreignKeyXID("user_id", 4),
		field.String("content").Annotations(
			entproto.Field(5),
		),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Field("user_id").
			Unique().Required().
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
				entgql.Skip(
					entgql.SkipMutationCreateInput,
					entgql.SkipMutationUpdateInput,
				),
				entproto.Skip(),
				entproto.Field(6),
			),
	}
}

func (Post) Indexes() []ent.Index {
	return nil
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
