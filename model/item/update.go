package item

import (
	"database/sql"
	"log"
)


func Update(db *sql.DB, id int64, name string) (Item, error)  {

	log.Println("item Update")
	defer log.Println("item Update end")

	item := Item{
		Id: id,
		Name: name,
	}

	query := `
		UPDATE t_user SET
			name = $2,
		    updated_at = now()
		WHERE
			id = $1
`

	_, err := db.Query(query, item.Id, item.Name)
	if err != nil{
		return item, err
	}

	return item, nil
}