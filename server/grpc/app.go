package main

import (
	"blueprint/db"
	"blueprint/ent/proto/entpb"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"entgo.io/ent/examples/fs/ent/migrate"
)

const defaultPort = "5000"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	client := db.Open()

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

	userService := entpb.NewUserService(client)
	postService := entpb.NewPostService(client)

	server := grpc.NewServer()

	entpb.RegisterUserServiceServer(server, userService)
	entpb.RegisterPostServiceServer(server, postService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	log.Printf("listening on :%s", port)
	// トラフィックを無期限にリッスン
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
