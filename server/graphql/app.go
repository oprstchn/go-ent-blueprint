package main

import (
	"blueprint/db"
	"blueprint/ent/migrate"
	"blueprint/graphql/resolver"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8081"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	client := db.MysqlOpen()

	// auto migration
	ctx := context.Background()
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		panic(err)
	}

	src := handler.NewDefaultServer(
		resolver.NewSchema(client),
	)

	src.Use(entgql.Transactioner{TxOpener: client})

	http.Handle("/", playground.Handler("Blueprint", "/query"))
	http.Handle("/query", src)

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
