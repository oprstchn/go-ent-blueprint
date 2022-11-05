// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	ent "blueprint/ent"
	user "blueprint/ent/user"
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	errors "errors"
	fmt "fmt"
	xid "github.com/rs/xid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// UserService implements UserServiceServer
type UserService struct {
	client *ent.Client
	UnimplementedUserServiceServer
}

// NewUserService returns a new UserService
func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		client: client,
	}
}

// toProtoUser transforms the ent type to the pb type
func toProtoUser(e *ent.User) (*User, error) {
	v := &User{}
	age := int64(e.Age)
	v.Age = age
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	idValue, err := e.ID.Value()
	if err != nil {
		return nil, err
	}
	id, ok := idValue.(string)
	if !ok {
		return nil, errors.New("casting value to string")
	}
	v.Id = id
	name := e.Name
	v.Name = name
	updated_at := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updated_at
	return v, nil
}

// toProtoUserList transforms a list of ent type to a list of pb type
func toProtoUserList(e []*ent.User) ([]*User, error) {
	var pbList []*User
	for _, entEntity := range e {
		pbEntity, err := toProtoUser(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements UserServiceServer.Create
func (svc *UserService) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	user := req.GetUser()
	m, err := svc.createBuilder(user)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUser(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	var (
		err error
		get *ent.User
	)
	id := xid.ID{}
	if err := (&id).Scan(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetUserRequest_VIEW_UNSPECIFIED, GetUserRequest_BASIC:
		get, err = svc.client.User.Get(ctx, id)
	case GetUserRequest_WITH_EDGE_IDS:
		get, err = svc.client.User.Query().
			Where(user.ID(id)).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUser(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements UserServiceServer.Update
func (svc *UserService) Update(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	user := req.GetUser()
	userID := xid.ID{}
	if err := (&userID).Scan(user.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.User.UpdateOneID(userID)
	userAge := int(user.GetAge())
	m.SetAge(userAge)
	userName := user.GetName()
	m.SetName(userName)
	userUpdatedAt := runtime.ExtractTime(user.GetUpdatedAt())
	m.SetUpdatedAt(userUpdatedAt)

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUser(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements UserServiceServer.Delete
func (svc *UserService) Delete(ctx context.Context, req *DeleteUserRequest) (*emptypb.Empty, error) {
	var err error
	id := xid.ID{}
	if err := (&id).Scan(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.User.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements UserServiceServer.List
func (svc *UserService) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	var (
		err      error
		entList  []*ent.User
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.User.Query().
		Order(ent.Desc(user.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken, err := xid.FromBytes(bytes)
		if err != nil {
    		return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		listQuery = listQuery.
			Where(user.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListUserRequest_VIEW_UNSPECIFIED, ListUserRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListUserRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoUserList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListUserResponse{
			UserList:      protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements UserServiceServer.BatchCreate
func (svc *UserService) BatchCreate(ctx context.Context, req *BatchCreateUsersRequest) (*BatchCreateUsersResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.UserCreate, len(requests))
	for i, req := range requests {
		user := req.GetUser()
		var err error
		bulk[i], err = svc.createBuilder(user)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.User.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoUserList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateUsersResponse{
			Users: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *UserService) createBuilder(user *User) (*ent.UserCreate, error) {
	m := svc.client.User.Create()
	userAge := int(user.GetAge())
	m.SetAge(userAge)
	userCreatedAt := runtime.ExtractTime(user.GetCreatedAt())
	m.SetCreatedAt(userCreatedAt)
	userName := user.GetName()
	m.SetName(userName)
	userUpdatedAt := runtime.ExtractTime(user.GetUpdatedAt())
	m.SetUpdatedAt(userUpdatedAt)
	return m, nil
}
