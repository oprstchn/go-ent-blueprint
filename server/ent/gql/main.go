package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graphql/schema/ent.graphql"),
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("../ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
			//gen.FeatureNamedEdges,
			//gen.FeatureSnapshot,
			//gen.FeatureSchemaConfig,
			//gen.FeatureLock,
			gen.FeatureModifier,
			gen.FeatureExecQuery,
			gen.FeatureUpsert,
			//gen.FeatureVersionedMigration,
		},
	}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
