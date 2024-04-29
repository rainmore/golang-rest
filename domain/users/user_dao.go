package users

import (
	"fmt"
	"time"

	"rainmore.com.au/rest-api/datasources/postgresql/users_db"
	"rainmore.com.au/rest-api/domain/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func Get(user *User) *errors.RestError {
	users_db.Ping()

	row := users_db.DBClient.QueryRow(
		"SELECT * FROM users WHERE id = $1", user.Id)

	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.CreatedAt); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func Save(user *User) *errors.RestError {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	var dateOfBirthStr string

	if user.DateOfBirth != nil {
		dateOfBirthStr = user.DateOfBirth.ToString()
	}

	row := users_db.DBClient.QueryRow(
		`INSERT INTO users (first_name, last_name, email, date_of_birth, created_at)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
		dateOfBirthStr,
		user.CreatedAt)

	if err := row.Scan(&user.Id); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return Get(user)
}
