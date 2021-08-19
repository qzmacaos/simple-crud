package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) UpdateItem(ctx context.Context, in *sea.UpdateItemRequest) (*sea.Item, error) {

	user := sea.Item{
		Id: in.GetId(),
		Name: in.GetName(),
	}

	query := `
		UPDATE t_user SET
			name = $2,
		    updated_at = now()
		WHERE
			id = $1
`

	_, err := s.DB.Query(query, user.Id, user.Name)
	if err != nil{
		return nil, err
	}

	return &user, nil
}
