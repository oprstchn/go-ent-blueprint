package db

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"blueprint/ent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func pgDatabaseURL() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		"root",
		"password",
		"0.0.0.0",
		"development",
	)
}

func mysqlDatabaseURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		"root",
		"password",
		"0.0.0.0",
		"development",
	)
}

func PgOpen() *ent.Client {
	db, err := sql.Open("pgx", pgDatabaseURL())
	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func MysqlOpen() *ent.Client {
	client, err := ent.Open("mysql", mysqlDatabaseURL())
	if err != nil {
		panic(err)
	}

	return client
}
