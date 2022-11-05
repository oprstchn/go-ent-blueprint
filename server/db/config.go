package db

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"blueprint/ent"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func databaseURL() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		"root",
		"password",
		"0.0.0.0",
		"development",
	)
}

func Open() *ent.Client {
	db, err := sql.Open("pgx", databaseURL())
	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
