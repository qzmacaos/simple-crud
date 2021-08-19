package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) ListUser(ctx context.Context, in *sea.ListUserRequest) (*sea.ListUserResponse, error){

	var(
		users []*sea.User
		pageFilter *sea.PageFilter
		offset,limit uint32
	)

	pageFilter = in.GetPageFilter()
	if pageFilter != nil{
		offset = pageFilter.GetPage() * s.CFG.Other.PageSize
		limit = pageFilter.GetLimit()
	}

	query := `
		SELECT id, name, age, user_type FROM 
		    t_user
		LIMIT $1
		OFFSET $2
`

	res, err := s.DB.Query(query, offset, limit)
	if err != nil{
		return nil, err
	}

	for res.Next(){
		user := sea.User{}
		res.Scan(
			&user.Id,
			&user.Name,
			&user.Age,
			&user.UserType)
		users = append(users, &user)
	}

	return &sea.ListUserResponse{Users: users}, nil
}
