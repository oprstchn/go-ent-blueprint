package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/rs/xid"
)

type BaseMixin struct {
	mixin.Schema
}

func xID(name string) ent.Field {
	return field.String(name).
		GoType(xid.ID{}).
		DefaultFunc(xid.New).
		Unique().
		NotEmpty().
		Immutable()
}

func foreignKeyXID(name string) ent.Field {
	return field.String(name).
		GoType(xid.ID{}).
		DefaultFunc(xid.New).
		Unique().
		NotEmpty()
}

// Fields of the User.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		xID("id"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entgql.Skip(
				entgql.SkipMutationCreateInput,
				entgql.SkipMutationUpdateInput,
			)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.Skip(
				entgql.SkipMutationCreateInput,
				entgql.SkipMutationUpdateInput,
			)),
	}
}

func (BaseMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
