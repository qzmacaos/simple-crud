package item

import (
	"database/sql"
	"errors"
	"log"
)


// Create - create item
func Create(db *sql.DB, name string,userId int64) (Item ,error) {

	log.Println("item Create")
	defer log.Println("item Create end")

	item := Item{
		Name: name,
		UserId: userId,
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

	res, err := db.Query(query, item.Name, item.UserId)
	if err != nil{
		return item, err
	}

	if res.Next() {
		if err = res.Scan(&item.Id); err != nil{
			return item, err
		}
	}else{
		return item, errors.New("no insert id")
	}


	return item, nil
}