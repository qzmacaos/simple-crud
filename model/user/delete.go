package user

import (
	"database/sql"
	"log"
)

// Delete - delete user
func Delete(db *sql.DB, id int64)error{

	log.Println("user Delete")
	defer log.Println("user Delete end")

	query := `
		DELETE FROM 
		    t_user 
		WHERE
			id = $1
`

	_, err := db.Query(query, id)
	if err != nil{
		return err
	}

	return nil
}