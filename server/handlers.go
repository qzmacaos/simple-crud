package server

import (
	"context"
	"simple-crud/model/item"
	"simple-crud/model/user"
	sea "simple-crud/serviceexampleapi"
	"strconv"
)

// CreateUser - proceed CreateUser request
func (s *Obj) CreateUser(ctx context.Context, in *sea.CreateUserRequest) (*sea.User, error) {
	var(
		name string
		age int32
		userType sea.UserType
	)
	name = in.GetName()
	age = in.GetAge()
	userType = in.GetUserType()

	user, err := user.Create(s.DB, name, age , userType)
	if err != nil{
		return nil, err
	}

	for _, item := range in.Items{
		s.CreateItem(ctx, item)
	}

	out := userToRpcUser(user)

	return &out, nil
}

// UpdateUser - proceed UpdateUser request
func (s *Obj) UpdateUser(ctx context.Context, in *sea.UpdateUserRequest) (*sea.User, error) {
	var(
		id int64
		name string
		age int32
		userType sea.UserType
	)

	id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil{
		return nil, err
	}

	name = in.GetName()
	age = in.GetAge()
	userType = in.GetUserType()


	user, err := user.Update(s.DB, id, name, age, userType)
	if err != nil{
		return nil, err
	}

	out := userToRpcUser(user)

	return &out, nil
}

// DeleteUser - proceed DeleteUser request
func (s *Obj) DeleteUser(ctx context.Context, in *sea.DeleteUserRequest) (*sea.DeleteUserResponse, error){


	var(
		id int64
	)

	id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil{
		return nil, err
	}

	if err := user.Delete(s.DB, id); err != nil {
		return nil, err
	}

	return &sea.DeleteUserResponse{}, nil
}

// GetUser - proceed GetUser request
func (s *Obj) GetUser(ctx context.Context, in *sea.GetUserRequest) (*sea.User, error) {
	var(
		id int64
	)

	id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil{
		return nil, err
	}

	user, err := user.Get(s.DB, id)
	if err != nil{
		return nil, err
	}

	out := userToRpcUser(user)

	return &out, nil
}

// ListUser - proceed ListUser request
func (s *Obj) ListUser(ctx context.Context, in *sea.ListUserRequest) (*sea.ListUserResponse, error){

	var(
		out []*sea.User
		pageFilter *sea.PageFilter
		offset,limit uint32
	)

	pageFilter = in.GetPageFilter()
	if pageFilter != nil{
		offset = pageFilter.GetPage() * pageSize
		limit = pageFilter.GetLimit()
	}

	users, err := user.List(s.DB, offset, limit)

	for _, user := range users{
		rpcUser := userToRpcUser(user)
		out = append(out, &rpcUser)
	}

	if err != nil {
		return nil ,err
	}

	return &sea.ListUserResponse{Users: out}, nil
}

// CreateItem - proceed CreateItem request
func (s *Obj) CreateItem(ctx context.Context, in *sea.CreateItemRequest) (*sea.Item, error) {

	var(
		name string
		userId int64
	)
	name = in.GetName()

	userId, err := strconv.ParseInt(in.GetUserId(), 10, 64)
	if err != nil{
		return nil, err
	}

	item, err := item.Create(s.DB, name ,userId)
	if err != nil{
		return nil, err
	}

	out := itemToRpcItem(item)

	return &out, nil
}

// UpdateItem - proceed UpdateItem request
func (s *Obj) UpdateItem(ctx context.Context, in *sea.UpdateItemRequest) (*sea.Item, error) {

	var(
		id int64
		name string
	)

	id, err := strconv.ParseInt(in.GetId(), 10, 64);
	if err != nil{
		return nil, err
	}

	name = in.GetName()

	item, err := item.Update(s.DB, id, name)
	if err != nil{
		return nil, err
	}

	out := itemToRpcItem(item)

	return &out, nil
}