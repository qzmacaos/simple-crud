package user

import (
	"database/sql"
	"log"
	sea "simple-crud/serviceexampleapi"
)


func Update(db *sql.DB, id int64, name string, age int32, userType sea.UserType) (User, error) {

	log.Println("user Update")
	defer log.Println("user Update end")

	user := User{
		Id: id,
		Name: name,
		Age: age,
		UserType: userType,
	}

	query := `
		UPDATE t_user SET
			name = $2,
			age = $3,
			user_type = $4,
			updated_at = now()
		WHERE
			id = $1
`

	_, err := db.Query(query, user.Id, user.Name, user.Age, user.UserType)
	if err != nil{
		return user, err
	}

	return user, nil
}