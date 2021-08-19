package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) CreateUser(ctx context.Context, in *sea.CreateUserRequest) (*sea.User, error) {

	user := sea.User{
		Name: in.GetName(),
		Age:in.GetAge(),
		UserType: in.GetUserType(),
	}

	query := `
		INSERT INTO t_user
			(
				name,
				age,
				user_type,
			 	updated_at,
			 	created_at
			)
		VALUES
		(
			$1,
			$2,
			$3,
			 now(),
			 now()
		)
		returning id
`

	res, err := s.DB.Query(query, user.Name, user.Age, user.UserType)
	if err != nil{
		return nil, err
	}

	if err = res.Scan(&user.Id); err != nil{
		return nil, err
	}

	for _, item := range in.Items{
		s.CreateItem(ctx, item)
	}

	return &user, nil

}
