package user

import (
	"database/sql"
	"errors"
	"log"
	"simple-crud/model/item"
)


func Get(db *sql.DB, id int64) (User, error) {

	log.Println("user Get")
	defer log.Println("user Get end")

	var(
		user User
	)

	user.Id = id
	query := `
		SELECT name, age, user_type FROM 
		    t_user 
		WHERE
			id = $1
`

	res, err := db.Query(query, user.Id)
	if err != nil{
		return user, err
	}

	if res.Next(){
		if err = res.Scan(&user.Name, &user.Age, &user.UserType); err != nil{
			return user, err
		}
	}else{
		return user, errors.New("no insert id")
	}

	userItems, err := item.ListItem(db, user.Id)
	if err != nil{
		return user, err
	}

	user.Items = userItems

	return user, nil
}