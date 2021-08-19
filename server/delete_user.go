package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) DeleteUser(ctx context.Context, in *sea.DeleteUserRequest) (*sea.DeleteUserResponse, error){

	var(
		id string
	)

	id = in.GetId()

	query := `
		DELETE FROM 
		    t_user 
		WHERE
			id = $1
`

	_, err := s.DB.Query(query, id)
	if err != nil{
		return nil, err
	}

	return nil, nil
}
