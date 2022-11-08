package ent

// //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature entql,sql/modifier,sql/lock,sql/upsert,sql/execquery,namedges ./schema

//for Graphql
//go:generate go run -mod=mod ./gql/main.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
//
//for gRPC
// //go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
