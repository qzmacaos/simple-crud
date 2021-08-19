package server

import sea "simple-crud/serviceexampleapi"

func (s *Obj)listItem(userId string)  ([]*sea.Item, error){

	var(
		items []*sea.Item
	)

	query := `
		SELECT 
			id, name
		FROM 
			t_item
		WHERE
			user_id = $1
`

	res, err := s.DB.Query(query, userId)
	if err != nil{
		return nil, err
	}

	for res.Next(){
		item := sea.Item{}
		res.Scan(&item.Id, &item.Name)
		items = append(items, &item)
	}
	return items, nil
}
