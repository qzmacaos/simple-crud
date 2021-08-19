package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) GetUser(ctx context.Context, in *sea.GetUserRequest) (*sea.User, error) {
	var(
		user sea.User
	)

	user.Id = in.GetId()

	query := `
		SELECT name, age, user_type FROM 
		    t_user 
		WHERE
			id = $1
`

	res, err := s.DB.Query(query, user.Id)
	if err != nil{
		return nil, err
	}

	if err = res.Scan(&user.Name, &user.Age, &user.UserType); err != nil{
		return nil, err
	}

	userItems, err := s.listItem(user.Id)
	if err != nil{
		return nil, err
	}

	user.Items = userItems

	return &user, nil
}
