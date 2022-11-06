package main

import (
	entconnect "blueprint/connect-grpc/ent/entpbconnect"
	"blueprint/db"
	"blueprint/ent/proto/entpb"
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"blueprint/ent/migrate"
)

const defaultPort = "5000"

type PostServer struct {
	service *entpb.PostService
}

func (p *PostServer) Create(
	ctx context.Context,
	req *connect.Request[entpb.CreatePostRequest],
) (*connect.Response[entpb.Post], error) {
	post, err := p.service.Create(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.Post]{
		Msg: post,
	}
	return res, nil
}

func (p *PostServer) Get(
	ctx context.Context,
	req *connect.Request[entpb.GetPostRequest],
) (*connect.Response[entpb.Post], error) {
	post, err := p.service.Get(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.Post]{
		Msg: post,
	}
	return res, nil
}

func (p *PostServer) Update(
	ctx context.Context,
	req *connect.Request[entpb.UpdatePostRequest],
) (*connect.Response[entpb.Post], error) {
	post, err := p.service.Update(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.Post]{
		Msg: post,
	}
	return res, nil
}

func (p *PostServer) Delete(
	ctx context.Context,
	req *connect.Request[entpb.DeletePostRequest],
) (*connect.Response[emptypb.Empty], error) {
	r, err := p.service.Delete(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[emptypb.Empty]{
		Msg: r,
	}
	return res, nil
}

func (p *PostServer) List(
	ctx context.Context,
	req *connect.Request[entpb.ListPostRequest],
) (*connect.Response[entpb.ListPostResponse], error) {
	posts, err := p.service.List(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.ListPostResponse]{
		Msg: posts,
	}
	return res, nil
}

func (p *PostServer) BatchCreate(
	ctx context.Context,
	req *connect.Request[entpb.BatchCreatePostsRequest],
) (*connect.Response[entpb.BatchCreatePostsResponse], error) {
	posts, err := p.service.BatchCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.BatchCreatePostsResponse]{
		Msg: posts,
	}
	return res, nil
}

type UserServer struct {
	service *entpb.UserService
}

func (p *UserServer) Create(
	ctx context.Context,
	req *connect.Request[entpb.CreateUserRequest],
) (*connect.Response[entpb.User], error) {
	user, err := p.service.Create(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.User]{
		Msg: user,
	}
	return res, nil
}

func (p *UserServer) Get(
	ctx context.Context,
	req *connect.Request[entpb.GetUserRequest],
) (*connect.Response[entpb.User], error) {
	user, err := p.service.Get(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.User]{
		Msg: user,
	}
	return res, nil
}

func (p *UserServer) Update(
	ctx context.Context,
	req *connect.Request[entpb.UpdateUserRequest],
) (*connect.Response[entpb.User], error) {
	user, err := p.service.Update(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.User]{
		Msg: user,
	}
	return res, nil
}

func (p *UserServer) Delete(
	ctx context.Context,
	req *connect.Request[entpb.DeleteUserRequest],
) (*connect.Response[emptypb.Empty], error) {
	r, err := p.service.Delete(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[emptypb.Empty]{
		Msg: r,
	}
	return res, nil
}

func (p *UserServer) List(
	ctx context.Context,
	req *connect.Request[entpb.ListUserRequest],
) (*connect.Response[entpb.ListUserResponse], error) {
	users, err := p.service.List(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.ListUserResponse]{
		Msg: users,
	}
	return res, nil
}

func (p *UserServer) BatchCreate(
	ctx context.Context,
	req *connect.Request[entpb.BatchCreateUsersRequest],
) (*connect.Response[entpb.BatchCreateUsersResponse], error) {
	users, err := p.service.BatchCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := &connect.Response[entpb.BatchCreateUsersResponse]{
		Msg: users,
	}
	return res, nil
}

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

	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(entconnect.NewPostServiceHandler(&PostServer{service: postService}))
	mux.Handle(entconnect.NewUserServiceHandler(&UserServer{service: userService}))
	err := http.ListenAndServe(
		fmt.Sprintf("localhost:%s", port),
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
