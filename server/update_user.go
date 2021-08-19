package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) UpdateUser(ctx context.Context, in *sea.UpdateUserRequest) (*sea.User, error) {

	user := sea.User{
		Id: in.GetId(),
		Name: in.GetName(),
		Age:in.GetAge(),
		UserType: in.GetUserType(),
	}

	query := `
		UPDATE t_user SET
			name = $2,
			age = $3,
			user_type = $4,
			updated_at = now()
		WHERE
			id = $1
`

	_, err := s.DB.Query(query, user.Id, user.Name, user.Age, user.UserType)
	if err != nil{
		return nil, err
	}

	return nil, nil
}
