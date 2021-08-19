package server

import (
	"simple-crud/model/item"
	"simple-crud/model/user"
	sea "simple-crud/serviceexampleapi"
	"strconv"
)

func itemToRpcItem(in item.Item) sea.Item {

	id := strconv.FormatInt(in.Id, 10)
	userId := strconv.FormatInt(in.UserId, 10)
	return sea.Item{
		Id:id,
		Name: in.Name,
		UserId: userId,
		UpdatedAt: in.UpdatedAt,
		CreatedAt: in.CreatedAt,
	}
}

func userToRpcUser(in user.User) sea.User{
	id := strconv.FormatInt(in.Id, 10)
	user := sea.User{
		Id:id,
		Name: in.Name,
		Age: in.Age,
		UserType: in.UserType,
		UpdatedAt: in.UpdatedAt,
		CreatedAt: in.CreatedAt,
	}

	for _, item := range in.Items{
		if item != nil{
			rpcItem := itemToRpcItem(*item)
			user.Items = append(user.Items, &rpcItem)
		}
	}
	return user
}