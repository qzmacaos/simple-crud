package item

import (
	"database/sql"
	"log"
)

// ListItem - get list of items attached to users
func ListItem(db *sql.DB, userId int64)  ([]*Item, error){

	log.Println("item ListItem")
	defer log.Println("item ListItem end")

	var(
		items []*Item
	)

	query := `
		SELECT 
			id, name
		FROM 
			t_item
		WHERE
			user_id = $1
`

	res, err := db.Query(query, userId)
	if err != nil{
		return nil, err
	}

	for res.Next(){
		item := Item{}
		res.Scan(&item.Id, &item.Name)
		items = append(items, &item)
	}
	return items, nil
}
