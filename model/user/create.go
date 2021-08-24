package user

import (
	"database/sql"
	"errors"
	"log"
	sea "simple-crud/serviceexampleapi"
)

// Create - create user
func Create(db *sql.DB, name string, age int32, userType sea.UserType) (User, error) {

	log.Println("user Create")
	defer log.Println("user Create end")

	user := User{
		Name: name,
		Age: age,
		UserType: userType,
	}

	query := `
		INSERT INTO t_user
			(
				name,
				age,
				user_type,
			 	updated_at,
			 	created_at
			)
		VALUES
		(
			$1,
			$2,
			$3,
			 now(),
			 now()
		)
		RETURNING id
`

	res, err := db.Query(query, user.Name, user.Age, user.UserType)
	if err != nil{
		return user, err
	}

	if res.Next(){
		if err = res.Scan(&user.Id); err != nil{
			return user, err
		}
	}else{
		return user, errors.New("no insert id")
	}

	return user, nil
}
