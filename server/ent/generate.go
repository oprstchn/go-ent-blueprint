package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./schema
//go:generate go run -mod=mod ./gql/main.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
