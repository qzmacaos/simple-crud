package server

import (
	"context"
	sea "simple-crud/serviceexampleapi"
)

func (s *Obj) CreateItem(ctx context.Context, in *sea.CreateItemRequest) (*sea.Item, error) {

	item := sea.Item{
		Name: in.GetName(),
		UserId: in.GetUserId(),
	}

	query := `
		INSERT INTO t_item
			(
				name,
				user_id,
				updated_at,
				created_at
			)
		VALUES
		(
			$1,
			$2,
			now(),
			now()
		)
		RETURNING id
`

	res, err := s.DB.Query(query, item.Name, item.UserId)
	if err != nil{
		return nil, err
	}


	if err = res.Scan(&item.Id); err != nil{
		return nil, err
	}

	return &item, nil

}
