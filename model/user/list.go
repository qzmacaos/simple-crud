package user

import (
	"database/sql"
	"log"
)

// List - get list of users with offset and limit
func List(db *sql.DB, offset, limit uint32) ([]User, error) {

	log.Println("user List")
	defer log.Println("user List end")

	var users []User

	query := `
		SELECT id, name, age, user_type FROM 
		    t_user
		LIMIT $1
		OFFSET $2
`
	log.Println(offset, limit)

	res, err := db.Query(query, limit, offset)
	if err != nil{
		return nil, err
	}

	for res.Next(){
		user := User{}
		if err = res.Scan(
			&user.Id,
			&user.Name,
			&user.Age,
			&user.UserType); err != nil{
				return nil, err
		}

		users = append(users, user)
	}
	log.Println(len(users))

	return users, nil
}
